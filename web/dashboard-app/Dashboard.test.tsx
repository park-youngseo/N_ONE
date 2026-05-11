import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import Dashboard from './Dashboard';

describe('Dashboard Component', () => {
  it('renders loading state initially', () => {
    render(<Dashboard />);
    expect(screen.getByText('로딩 중...')).toBeInTheDocument();
  });

  it('renders student name and attendance rate after loading', async () => {
    render(<Dashboard />);
    
    // Wait for the mock API call to resolve
    await waitFor(() => {
      expect(screen.getByText('홍길동 님, 환영합니다!')).toBeInTheDocument();
    }, { timeout: 1500 });

    expect(screen.getByText('83%')).toBeInTheDocument();
    expect(screen.getByText('로미오와 줄리엣')).toBeInTheDocument();
    expect(screen.getByText('50% 완료')).toBeInTheDocument();
  });

  it('toggles theme between light and dark', async () => {
    render(<Dashboard />);
    
    await waitFor(() => {
      expect(screen.getByText('홍길동 님, 환영합니다!')).toBeInTheDocument();
    }, { timeout: 1500 });

    const toggleButton = screen.getByRole('button', { name: /다크 모드/i });
    const container = screen.getByText('N-ONE 연기학원').closest('.dashboard-container');
    
    expect(container).toHaveClass('light');
    
    fireEvent.click(toggleButton);
    
    expect(container).toHaveClass('dark');
    expect(screen.getByRole('button', { name: /라이트 모드/i })).toBeInTheDocument();
  });
});
