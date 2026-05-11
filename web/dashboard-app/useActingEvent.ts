import { useEffect, useRef, useState, useCallback } from 'react';

export interface ScriptProgressEvent {
  script_id: string;
  scene_id: string;
  progress_percent: number;
  timestamp: number;
}

interface UseActingEventReturn {
  isConnected: boolean;
  latestEvent: ScriptProgressEvent | null;
  publishOptimisticEvent: (event: ScriptProgressEvent) => void;
}

export const useActingEvent = (url: string): UseActingEventReturn => {
  const [isConnected, setIsConnected] = useState(false);
  const [latestEvent, setLatestEvent] = useState<ScriptProgressEvent | null>(null);
  
  const wsRef = useRef<WebSocket | null>(null);
  const reconnectTimeoutRef = useRef<NodeJS.Timeout | null>(null);
  const reconnectCountRef = useRef(0);
  
  // Offline Event Queue for Exponential Backoff / Reconnect scenario
  const offlineQueueRef = useRef<ScriptProgressEvent[]>([]);

  const connect = useCallback(() => {
    if (wsRef.current?.readyState === WebSocket.OPEN) return;

    console.log(`[WebSocket] Connecting to ${url}...`);
    const ws = new WebSocket(url);
    wsRef.current = ws;

    ws.onopen = () => {
      console.log('[WebSocket] Connected successfully');
      setIsConnected(true);
      reconnectCountRef.current = 0;

      // Flush offline queue when connected
      if (offlineQueueRef.current.length > 0) {
        console.log(`[WebSocket] Flushing ${offlineQueueRef.current.length} offline events...`);
        // In a real implementation, you would send this to the backend REST API
        // or through the WebSocket channel if supported.
        // For this scaffold, we just clear the queue to simulate successful sync.
        offlineQueueRef.current = [];
      }
    };

    ws.onmessage = (messageEvent) => {
      try {
        const data: ScriptProgressEvent = JSON.parse(messageEvent.data);
        console.log('[WebSocket] Received event:', data);
        setLatestEvent(data);
      } catch (e) {
        console.error('[WebSocket] Failed to parse message:', e);
      }
    };

    ws.onclose = (event) => {
      console.log(`[WebSocket] Disconnected: code=${event.code}, reason=${event.reason}`);
      setIsConnected(false);
      wsRef.current = null;
      
      // Exponential Backoff Reconnect Logic
      const backoffDelay = Math.min(1000 * Math.pow(2, reconnectCountRef.current), 30000);
      console.log(`[WebSocket] Scheduling reconnect in ${backoffDelay}ms...`);
      
      if (reconnectTimeoutRef.current) clearTimeout(reconnectTimeoutRef.current);
      reconnectTimeoutRef.current = setTimeout(() => {
        reconnectCountRef.current++;
        connect();
      }, backoffDelay);
    };

    ws.onerror = (error) => {
      console.error('[WebSocket] Error occurred:', error);
      // onclose will be called automatically after onerror
    };
  }, [url]);

  useEffect(() => {
    connect();

    return () => {
      if (reconnectTimeoutRef.current) clearTimeout(reconnectTimeoutRef.current);
      if (wsRef.current) {
        wsRef.current.close();
      }
    };
  }, [connect]);

  // Method to allow UI to trigger optimistic updates, queuing them if offline
  const publishOptimisticEvent = useCallback((event: ScriptProgressEvent) => {
    if (isConnected) {
      // In a full implementation, call REST API here to persist,
      // which triggers the Redis Pub/Sub back to us.
      console.log('[Action] Publishing event online:', event);
      // Optimistically update local state immediately
      setLatestEvent(event);
    } else {
      console.log('[Action] Offline. Queuing event:', event);
      offlineQueueRef.current.push(event);
      // Optimistically update local state even when offline
      setLatestEvent(event);
    }
  }, [isConnected]);

  return { isConnected, latestEvent, publishOptimisticEvent };
};
