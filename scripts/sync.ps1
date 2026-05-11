# Autopus-ADK Auto-Vault Sync Script (Protocol-Strict Version)

param (
    [string]$Message = "Auto-sync: $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')"
)

[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

Write-Host "`n🐙 Autopus 물류 시스템 가동 중..." -ForegroundColor Cyan

git add .
git reset "제미나이,GPT API키.txt" 2>$null

$msgFile = "$env:TEMP\autopus_commit_msg.txt"
$utf8NoBom = New-Object System.Text.UTF8Encoding($false)

# 메시지 조립 (Lore 정규 규격: 헤더-빈줄-본문-빈줄-사인오프)
$header = "feat(sync): $Message"
$body = "Constraint: AX-ALIGNMENT-300"
$signoff = "🐙 Autopus <noreply@autopus.co>"

# 각 섹션 사이에 명확한 빈 줄을 삽입하고 마지막에 뉴라인 추가
$fullMsg = "$header`n`n$body`n`n$signoff`n"

[System.IO.File]::WriteAllText($msgFile, $fullMsg, $utf8NoBom)

Write-Host "📦 로컬 금고에 기록 중..." -ForegroundColor Yellow
git commit -F $msgFile

Write-Host "🚚 중앙 창고(GitHub)로 배달 중..." -ForegroundColor Green
& "C:\Program Files\GitHub CLI\gh.exe" auth setup-git
git push origin master

if (Test-Path $msgFile) { Remove-Item $msgFile }

Write-Host "✅ 배달 완료! 모든 지식이 안전하게 보관되었습니다.`n" -ForegroundColor Green
