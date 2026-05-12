# [Master 10-02] AI Director 및 메타휴먼 지능형 연출 전략

사령관님, 이제 언리얼 엔진은 사령관님의 지시를 기다리는 **'자율 주행 영화 스튜디오'**가 될 것입니다.

## 1. AI Director (Python Automation)
- **핵심**: 언리얼 엔진의 Python API를 활용하여 LLM이 직접 씬(Scene)을 구성함.
- **실전**: 
    - AI가 "숲속의 추격전" 시나리오를 읽고, **PCG(00-01)** 시스템을 호출하여 숲을 만들고, 카메라를 액션 구도로 자동 배치.
    - **Blueprint Automation**: 반복적인 블루프린트 노드 구성을 AI가 직접 생성하여 로직 구축 시간 단축.

## 2. MetaHuman AI Performance (Audio-to-Anim)
- **Facial Animation**: AI 음성 파일(15번 결과물)을 입력하면, `Audio2Face` 기술을 통해 메타휴먼의 입 모양, 눈 떨림, 표정 변화를 실시간 생성.
- **Body Motion**: 시나리오의 감정 태그(슬픔, 분노 등)를 읽고, AI가 가장 적절한 모션 캡처 데이터를 라이브러리에서 선택하여 캐릭터에 입힘.

---

## 3. NONE Scenario-to-Visual 파이프라인
- **Metadata Linking**: 01번 시나리오 엔진이 생성하는 각 장면에 **[Mood, Camera, Lighting]** 메타데이터를 포함시킴.
- **Auto-Rendering**: 언리얼 엔진이 이 메타데이터를 수신하여 자동으로 세트를 세팅하고, 고퀄리티 렌더링(Movie Render Queue)을 실행하여 최종 영상 파일 생성.

---
*이 지식은 사령관님의 시나리오에 '육체'와 '표정'을 부여하여, 상상을 현실의 이미지로 고착시키는 최종 병기입니다.*