import React, { useState, useEffect } from 'react';
import './Dashboard.css';

// Types mimicking the Go backend
interface Student {
  id: string;
  name: string;
}

interface Script {
  id: string;
  title: string;
  total_scenes: number;
}

interface DashboardSummary {
  student: Student;
  attendance_rate: number;
  recent_script?: Script;
}

const Dashboard: React.FC = () => {
  const [summary, setSummary] = useState<DashboardSummary | null>(null);
  const [loading, setLoading] = useState(true);
  const [theme, setTheme] = useState<'light' | 'dark'>('light');

  useEffect(() => {
    // Mock API Call
    setTimeout(() => {
      setSummary({
        student: { id: 'stu-1', name: '홍길동' },
        attendance_rate: 83, // Mock data based on acceptance criteria
        recent_script: { id: 'script-1', title: '로미오와 줄리엣', total_scenes: 2 }
      });
      setLoading(false);
    }, 1000);
  }, []);

  const toggleTheme = () => {
    setTheme(prev => prev === 'light' ? 'dark' : 'light');
  };

  if (loading) {
    return <div className={`dashboard-container ${theme}`}><div className="spinner">로딩 중...</div></div>;
  }

  return (
    <div className={`dashboard-container ${theme}`}>
      <header className="dashboard-header">
        <h1>N-ONE 연기학원</h1>
        <div className="header-actions">
          <span>{summary?.student.name} 님, 환영합니다!</span>
          <button onClick={toggleTheme}>
            {theme === 'light' ? '🌙 다크 모드' : '☀️ 라이트 모드'}
          </button>
        </div>
      </header>

      <main className="dashboard-main">
        <section className="card attendance-card">
          <h2>이번 달 출결 현황</h2>
          <div className="attendance-rate">
            <span className="rate-value">{summary?.attendance_rate}%</span>
            <span className="rate-label">출석률</span>
          </div>
          <div className="calendar-placeholder">
            {/* Calendar component would go here */}
            [ 달력 위젯 (Mock) ]
          </div>
        </section>

        <section className="card script-card">
          <h2>최근 대본 연습</h2>
          {summary?.recent_script ? (
            <div className="script-progress-container">
              <h3>{summary.recent_script.title}</h3>
              <div className="progress-bar-bg">
                <div 
                  className="progress-bar-fill" 
                  style={{ width: '50%' }} // Mock 50%
                ></div>
              </div>
              <p className="progress-text">50% 완료</p>
              <button className="btn-primary">이어서 연습하기</button>
            </div>
          ) : (
            <p className="empty-state">진행 중인 대본이 없습니다.</p>
          )}
        </section>
      </main>
    </div>
  );
};

export default Dashboard;
