# N-ONE 올라마 연동 테스트 스크립트
# 올라마가 정상 가동 중인지, 하네스와의 통신이 되는지 검증한다.

[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

Write-Host "`n🐙 N-ONE 올라마 연동 테스트 시작..." -ForegroundColor Cyan

# 1. 올라마 서버 가동 확인 (Ping)
Write-Host "`n[1/4] 올라마 서버 상태 확인..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://127.0.0.1:11434/api/tags" -Method GET -TimeoutSec 5
    $models = $response.models | ForEach-Object { $_.name }
    Write-Host "  ✅ 올라마 서버 가동 중. 사용 가능 모델: $($models -join ', ')" -ForegroundColor Green
} catch {
    Write-Host "  ❌ 올라마 서버 미응답. 'ollama serve'를 먼저 실행하세요." -ForegroundColor Red
    exit 1
}

# 2. qwen2.5:14b 모델 존재 확인
Write-Host "`n[2/4] 주력 모델(qwen2.5:14b) 확인..." -ForegroundColor Yellow
if ($models -contains "qwen2.5:14b") {
    Write-Host "  ✅ qwen2.5:14b 모델 장착 완료" -ForegroundColor Green
} else {
    Write-Host "  ⚠️  qwen2.5:14b 미발견. 'ollama pull qwen2.5:14b'로 설치하세요." -ForegroundColor Yellow
}

# 3. 실제 추론 테스트 (단순 전처리 작업)
Write-Host "`n[3/4] 단순 전처리 작업 테스트 (JSON 포맷팅)..." -ForegroundColor Yellow
$body = @{
    model = "qwen2.5:14b"
    prompt = "다음 데이터를 JSON 형식으로 변환해라: name=Park, age=30, city=Seoul"
    system = "너는 단순 전처리 작업자다. 창의성을 발휘하지 말고 주어진 원본 데이터를 지시대로만 변환하여 출력해라."
    stream = $false
    options = @{ temperature = 0.0 }
} | ConvertTo-Json

try {
    $result = Invoke-RestMethod -Uri "http://127.0.0.1:11434/api/generate" -Method POST -Body $body -ContentType "application/json" -TimeoutSec 60
    Write-Host "  ✅ 추론 성공! 응답:" -ForegroundColor Green
    Write-Host "  $($result.response.Substring(0, [Math]::Min(200, $result.response.Length)))" -ForegroundColor White
} catch {
    Write-Host "  ❌ 추론 실패: $_" -ForegroundColor Red
}

# 4. 하네스 엔진 빌드 확인
Write-Host "`n[4/4] 하네스 엔진 빌드 확인..." -ForegroundColor Yellow
go build ./cmd/auto 2>&1 | Out-Null
if ($LASTEXITCODE -eq 0) {
    Write-Host "  ✅ 하네스 엔진 빌드 성공 (올라마 어댑터 포함)" -ForegroundColor Green
} else {
    Write-Host "  ❌ 하네스 엔진 빌드 실패. 어댑터 코드 점검 필요." -ForegroundColor Red
}

Write-Host "`n🐙 테스트 완료!`n" -ForegroundColor Cyan
