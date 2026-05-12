25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **바로 되는 해법 1 (가장 쉽고 효과 큼)** 

**Script.bat 다시 실행 → “비공개/연령제한 영상 있나요?”에** `Y` **→** `chrome` **입력** → 내 스크립트가 **브라우저 쿠키** 를 써서 요청을 처리해. 이렇게 하면 **차단이 훨씬 덜 걸림** . 

## **바로 되는 해법 2 (요청 사이에 잠깐 쉬어가기)** 

지금 쓰는 `Youtube.txt.py` 에 **딱 몇 줄** 만 추가해서, 영상/재생목록마다 **10~25초 랜덤 대 기** 를 넣자. 

(방식은 안 바꾼다. 파일명/흐름 그대로, “한 줄 추가” 수준) 

## **1) 맨 위 import 줄에 아래 2개를 “추가”** 

python 코드 복사 

```
import time, random
```

## **2)** `main()` **또는 URL을 처리하는 반복문에 “대기” 넣기** 

`for u in urls:` 바로 안쪽에 **이 3줄** 추가: python 코드 복사 `delay = random.randint(10, 25)` _`#`_ `요청 간 랜덤 대기` _`(`_ `유튜브 차단 회피` _`)`_ `print(f"[대기] {delay}초 후 시작: {u}") time.sleep(delay)` 예시 (네 엔진 구조 기준): python 코드 복사 `for u in urls: delay = random.randint(10, 25) print(f"[대기] {delay}초 후 시작: {u}") time.sleep(delay) print(f"[처리] {u}") items = download_texts_from_url(u, tmp_dir, args.lang, args.cookies_from_brow ...` 

## 예시 (네 엔진 구조 기준): 

  

## **(선택) 재생목록 안에서도 조금 쉬어가기** 

`download_texts_from_url()` 에서 `for it in items:` 안쪽 **가장 끝** 에 한 줄 추가: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

310/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

코드 복사 

python 

```
time.sleep(random.randint(3, 8))  # 같은재생목록내영상사이간격
```

이 정도만 넣어도 **차단 빈도 확 내려감** . (메시지에 “-t sleep” 비슷한 권장과 같은 효과) 

## **추가로 알아두면 좋아** 

- **urls.txt를 잘게 나눠서** (예: 20~30개 단위) 여러 번 돌리면 더 안전함. **동시에 여러 창으로 돌리지 않기.** 한 창만. 

- 특정 영상이 계속 막히면 잠시 **주석(#)** 처리하고 나머지부터 진행. 

## **지금 당장 할 것 (3줄 요약)** 

**1. 가장 쉬운 방법:** Script.bat 다시 실행 → `Y` → `chrome` . 

**2.** 그래도 막히면, **위 “대기” 3줄** 을 `Youtube.txt.py` 의 `for u in urls:` 바로 안에 붙여 넣고 저장. 

**3.** 다시 Script.bat 실행. 

다시 돌려서 또 메시지가 뜨면, **마지막 3~4줄만** 붙여줘. 거기서 바로 이어서 잡을게. 

**==> picture [489 x 39] intentionally omitted <==**

C:\Users\엔오원\Desktop\yotube_script>import time, random 

'import'은(는) 내부 또는 외부 명령, 실행할 수 있는 프로 그램, 또는 

배치 파일이 아닙니다. 야.. 씨발 진짜 어디다 경로를 넣어 야 하는지 알려주라고 씨발놈아 니 멋대로 계속 진행 하지 말고 

## 형, 방금 그건 **CMD(검은창)에** `import` **를 쳤기 때문** 이야. 

`import time, random` 같은 건 **파이썬 파일 안** (= `Youtube.txt.py` 내용) 에 넣어야 해. 지금부터 **어디를 어떻게 고칠지 “정확한 위치”**를 알려줄게. 그대로 따라 해. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

311/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **딱 이대로 하자** 

## **1) 파이썬 파일을 메모장으로 연다** 

아래 줄을 **그대로** 복사해서 검은창에 붙여넣고 Enter: 

bat 코드 복사 

```
notepad "C:\Users\엔오원\Desktop\yotube_script\Youtube.txt.py"
```

## **2) 맨 위에 “대기 기능” 임포트 추가** 

메모장이 열리면 **파일 맨 위쪽** 에 보통 이런 줄이 있어: 

python 코드 복사 

```
import argparse, re, html, shutil
```

이 줄 또는 그 근처에 **아래 한 줄을 추가** 해: 

python 코드 복사 

```
import time, random
```

(같은 줄에 `, time, random` 을 붙여도 되고, 새 줄로 써도 됨) 

## **3) URL 하나 처리 시작 전에 “랜덤 대기” 넣기** 

## 메모장에서 **아래 코드 블록** 을 찾아: 

python 코드 복사 

```
for u in urls:
```

```
print(f"[처리] {u}")
```

`items = download_texts_from_url(u, tmp_dir, args.lang, args.cookies_from_brow for _, _, t in items: all_parts.append(t); all_parts.append("\n\n")`   

## 여기서 `for u in urls:` **바로 아래** 에 이 **3줄** 을 끼워 넣어줘: 

python 코드 복사 

```
    delay = random.randint(10, 25)  # 요청간랜덤대기(차단회피)
print(f"[대기] {delay}초후시작: {u}")
    time.sleep(delay)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 312/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## 수정 후 모양이 이렇게 되면 정답: 

python 

코드 복사 

```
for u in urls:
    delay = random.randint(10, 25)  # 요청간랜덤대기(차단회피)
print(f"[대기] {delay}초후시작: {u}")
    time.sleep(delay)
```

```
print(f"[처리] {u}")
    items = download_texts_from_url(u, tmp_dir, args.lang, args.cookies_from_brow
for _, _, t in items:
```

```
        all_parts.append(t); all_parts.append("\n\n")
```

  

## **4) (선택) 같은 재생목록 안에서도 잠깐 쉬자** 

메모장에서 `download_texts_from_url` **함수** 로 내려가서, 그 안의 

python 코드 복사 `for it in items: ...` 

## 루프 **맨 끝** 에 **이 한 줄** 을 넣어줘: 

python 코드 복사 

```
    time.sleep(random.randint(3, 8))  # 같은재생목록내부대기
```

## 예시: 

python 코드 복사 `for it in items: if not it: continue ... if txt: results.append((vid, title, txt)) time.sleep(random.randint(3, 8))` _`#`_ `같은 재생목록 내부 대기` 

## **5) 저장하고 닫기** 

메모장에서 **Ctrl + S(저장)** → 창 **닫기** . 

## **6) 실행 (쿠키도 함께 쓰면 더 안전)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

313/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

이제 다시 **배치** 로 돌리면 돼: 

**1.** `Script.bat` 더블클릭 

**2.** 정리 폴더 경로 입력(엔터=현재 폴더) 

**3.** 메모장에 링크 붙여넣고 **저장 → 닫기** 

**4.** “비공개/연령제한 있나요?” → **Y** 입력 후 **chrome** (엣지면 edge) 입력 권장 → 이렇게 하면 **브라우저 쿠키** 로 로그인 세션을 써서 **차단 덜 걸림** 

**5.** 처리 시작 

## **만약 여전히 막히면** 

검은창에서 **마지막 3~4줄** (에러 메시지)만 복사해서 보내줘. 그 줄 기준으로 바로 다음 수정점 딱 집어서 알려줄게. 

**==> picture [489 x 38] intentionally omitted <==**

아무것도 진행 사항이 되지 않아 멈춰져 있어 그냥 파이썬 이 

13s 동안 생각함 

형, 멈춘 이유는 보통 **(A) 메모장 안 닫음** , **(B) 대기(sleep) 중** , **(C) 질문 대기(Y/N) 중** , **(D) 링크/권한 문제** 중 하나야. 딱 필요한 것만 순서대로 잡자. 

## **0) 지금 당장 확인 (30초)** 

## **1. 메모장(urls.txt)이 아직 열려 있나?** 

→ 열려 있으면 **Ctrl+S(저장) → 닫기** . 닫아야 다음 단계가 진행돼. 

**2.** 검은창에 ** `[질문] 비공개/연령제한… (Y/N)` **이 떠 있는데 입력 안 했나? → **N** 치고 Enter (모르면 N). 

**3.** 검은창에 `[대기] 10~25초 후 시작:` 같은 줄이 보이면 **대기 중** 이야. (링크가 많으면 몇 분 보일 수 있음) 

## **1) 멈춤 원인별 ‘한 방’ 해결** 

## **A) 대기(sleep) 없애고 “즉시 처리”로 바꾸기** 

**1.** `Youtube.txt.py` 를 메모장으로 열기: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

314/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

코드 복사 

nginx 

 

```
notepad"C:\Users\엔오원\Desktop\yotube_script\Youtube.txt.py"
```

**2.** 아래 두 군데를 **그대로** 고쳐 저장. 

   - (검색) `for u in urls:` 바로 다음 줄들에서 

python 코드 복사 

```
delay = random.randint(10, 25)
print(f"[대기] {delay}초후시작: {u}")
time.sleep(delay)
```

## 를 **아래처럼 바꾸기** : 

python 코드 복사 

```
delay = 0
print(f"[대기] {delay}초후시작: {u}")
# time.sleep(delay)
```

(검색) `for it in items:` 루프 끝에서 python 코드 복사 `time.sleep(random.randint(3, 8))` 을 **주석 처리** : python 코드 복사 