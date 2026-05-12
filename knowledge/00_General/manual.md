# 🐙 Marketing Automation System User Manual

본 문서는 `projects/NONE_마케팅/src/`에 구현된 자동화 시스템의 사용 방법을 안내합니다.

## 1. 시스템 실행 방법
터미널에서 아래 명령어를 실행하면 통합 파이프라인이 가동됩니다.
```powershell
python projects/NONE_마케팅/src/main_automation.py
```
**실행 순서**: 지식 로드 → 콘텐츠 생성 → 예약 배포 → 성과 추적

## 2. 주요 모듈 설명
- **Knowledge Processor**: `knowledge/` 폴더의 데이터를 분석하여 인사이트 추출.
- **Content Generator**: 지식을 바탕으로 인스타 릴스 대본 및 블로그 포스팅 생성.
- **Distributor Engine**: 4-1-1 법칙에 따른 채널별 예약 배포 관리.
- **Engagement Bot**: 유입 클릭수 및 성과 분석(CTR).

## 3. 커스터마이징 (Customization)
### 주제 변경
`main_automation.py` 파일의 `topic` 변수를 원하는 주제로 수정하세요.
```python
topic = "원하는 주제 입력"
```

### 지식 추가
`projects/NONE_마케팅/knowledge/` 폴더에 새로운 `.md` 파일을 추가하면 AI가 다음 실행 시 해당 지식을 자동으로 학습합니다.

## 4. 실제 API 연동 가이드
현재 시스템은 시뮬레이션 모드입니다. 실제 포스팅을 원하시면 다음 파일을 수정하세요:
1. **Instagram**: `distributor_engine.py` 내의 `schedule_post` 함수에 Facebook Graph API 연동 로직 추가.
2. **Blog**: 네이버/티스토리 API Key를 발급받아 연동.

---
*Created by Autopus-ADK (Gravity Assistant)*