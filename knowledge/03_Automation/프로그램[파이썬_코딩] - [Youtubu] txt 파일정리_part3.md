
**==> picture [193 x 218] intentionally omitted <==**

## 이렇게 나온다 

42s 동안 생각함 

형, 스샷처럼 **.bak** 가 생긴 건 내가 v1.6에서 안전 때문에 백업을 자동으로 만들게 해놨었 기 때문이야. 

요청대로 이제 **백업 안 만들고** , **원본이름.txt 그대로 덮어쓰기만** 하도록 고쳤어. 최신본 받기: Youtube_script.py (v1.7) 

지금 동작(간단 요약) 

`Kind:` · `Language:` 제거 → 리터럴 줄바꿈 정리 

- 종결어(요/다) 기준 문장 분리 → 각 문장 끝에 마침표 강제 문장 2개씩 한 줄 

## **백업 없음** , **파일명 접미사 없음** , **그대로 덮어쓰기** 

폴더에 이미 생겨있는 `.bak` 파일들은 필요 없으면 지워도 돼. 다음에 바꿀 규칙(예: 3문장 묶음, 종결어 추가 등) 있으면 바로 말해줘. 

## txt 파일 Roger Kerin 마케팅강의 7장 

[5LUkDDF9hI0].ko.vtt.clean 처럼 출력하지 말고 Roger Kerin 마케팅강의 7장.txt 이렇게도 만들어 달라고 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/690091b4-9dc8-8322-b248-ec9fa57cfb21 

12/13 

25. 11. 30. 오후 1:04 

프로그램[파이썬/코딩] - [Youtubu] txt 파일정리 

35s 동안 생각함 

 

## 좋아 형. 요청대로 **파일명 정리 규칙** 까지 넣은 **v1.8** 준비 완료야. 

이제부터 입력 파일이 아래처럼 생겼어도: 

```
Roger Kerin 마케팅강의 7장 [5LUkDDF9hI0].ko.vtt.clean
```

   - `… .ko.vtt.clean` , `… .vtt.clean` , `… .vtt` , `… .clean` , `.txt` 등 

- 출력은 항상 `Roger Kerin 마케팅강의 7장.txt` 형식으로 만들어줘. 

(끝의 `[유튜브ID]` 블록과 모든 접미사/확장자는 제거 → `.txt` 로 저장) 

받기: 

최신 스크립트: Youtube_script.py (v1.8) 

사용법(초간단): 

**1.** `Youtube_script.py` 와 같은 폴더에 원본 파일들( `*.txt` , `*.clean` , `*.vtt` , `*.vtt.clean` )을 넣고 

**2.** 더블클릭(또는 `python Youtube_script.py` ) 

**3.** 결과가 **단순 제목형 `.txt` **로 생성/덮어쓰기됨 

추가로 바꿔줄 수 있어: 

- 파일명에 남기고 싶은 단어 화이트리스트(예: “강의 7장” 유지) ID가 중간에 있을 때(끝이 아닌)도 제거 

- 결과를 별도 폴더( `/output` )에만 저장 

다음 샘플 파일로 바로 돌려보자—문제 있으면 바로 손볼게. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/690091b4-9dc8-8322-b248-ec9fa57cfb21 

13/13 
