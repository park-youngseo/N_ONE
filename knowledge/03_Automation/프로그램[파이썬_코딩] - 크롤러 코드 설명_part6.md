
형, 좋아. 지금 화면 기준으로 **“첫 시작”이 딱 두 군데** 야: 

**1. 가운데 파란 버튼** `Open Folder` 

**2. 오른쪽 위 ‘Agent’ 패널의 채팅 입력칸(** `Ask anything` **)** 

이 두 군데를 순서대로 쓰면 된다. 

완전 처음부터 하나씩 정리해 줄게. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

27/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

## **1단계. 폴더부터 열기 (Open Folder)** 

**1.** 윈도우에서 먼저 

바탕화면에 `n1_crawler` 라는 폴더 하나 만들어 둔다. 

(이미 만들었으면 패스) 

**2.** 지금 보여준 Antigravity 화면에서 

가운데 큰 파란 버튼 `Open Folder` 클릭. 

**3.** 폴더 선택 창이 뜨면 

바탕화면 → `n1_crawler` 폴더 선택 → **폴더 선택** 클릭. 

**4.** 그러면 왼쪽에 파일 탐색기(Explorer)가 열리고, 

   - 상단 제목에 `n1_crawler` 이런 식으로 현재 폴더가 표시될 거야. 

   - → 여기까지가 **프로젝트 “첫 시작 1번”** 이라고 보면 돼. 

## **2단계. 에이전트에게 프롬프트 넣는 곳 (오른쪽 Agent 패널)** 

형이 물어본 “프롬프트 첫 입력”은 **오른쪽 Agent 패널** 이야. 

**1.** 화면 오른쪽에 `Agent` 라고 써 있고, 

   - 그 아래에 `Ask anything (Ctrl+L)` 라고 적힌 입력칸이 보일 거야. 

**2.** 그 입력칸을 마우스로 한 번 클릭한다. 

**3.** 내가 전에 만들어 준 **그래비티용 긴 프롬프트** (규칙 + v2.4 코드 전체) 를 

   - **그 입력칸에 그대로 붙여넣기** 한다. 

**4.** 붙여넣고 나서 

엔터 치거나, 

- 입력칸 오른쪽에 있는 ➜(전송 버튼) 클릭. 

## 이게 **“그래비티에 이 파일을 넣는 첫 시작 2번”** 이다. 

그 다음부터는: 

- 에이전트가 `cleaner.py` , `tagger.py` … 같은 코드들을 쫙 뱉어 줄 거고, 형은 왼쪽 파일 탐색기에서 

- `새 파일(New File)` 만들어서 → 이름 맞춰 저장 → 코드 붙여넣기만 하면 돼. 

## **3단계. Agent Manager는 나중 단계** 

가운데의 `Open Agent Manager` 버튼은 

- “전용 에이전트를 따로 저장해 두고 나중에도 계속 쓰고 싶다” 할 때 쓰는 거고, 

- 지금은 **그냥 오른쪽 Agent 패널** 에서 시작하는 게 더 간단하다. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

28/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

정리하면: 

**1. 첫 시작 1:** 가운데 `Open Folder` 눌러서 `n1_crawler` 폴더 연다. 

**2. 첫 시작 2:** 오른쪽 `Agent` **–** `Ask anything` **칸** 에 

   - v2.4 전체 프롬프트를 붙여넣고 전송한다. 

여기까지 따라 하고 나서, 

에이전트가 뱉어준 코드로 **어떤 파일부터 만들지** 단계별로 이어서 정리해 줄게. 

**==> picture [489 x 38] intentionally omitted <==**

**==> picture [289 x 180] intentionally omitted <==**

이건 뭐냐? 

## 형, 이건 **“이 폴더 안에 있는 코드 믿을 수 있냐?”** 물어보는 보안 창이야. 

- Antigravity(=VS Code 기반)에서 

- 폴더 안에 있는 코드/스크립트를 **AI가 자동으로 실행하거나 분석** 할 수 있어서, 혹시 악성 코드면 위험하니까 한 번 더 확인하는 단계야. 

지금 폴더 경로가 

## `C:\Users\psm23\Desktop\n1_crawler` 

→ 형이 직접 만든 폴더니까 **형이 작성자 = 믿을 수 있는 코드** 인 거지. 

## **형이 지금 해야 할 것** 

## **1.** 아래 체크 박스: 

```
Trust the authors of all files in the parent folder 'Desktop'
```

이건 바탕화면 전체를 신뢰할 거냐는 옵션이라, 

부담되면 체크 안 해도 되고, 형 컴퓨터 혼자 쓰면 체크해도 괜찮아. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

29/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

**2.** 파란 버튼 `Yes, I trust the authors` 클릭 

   - → 그래야 **Restricted Mode(제한 모드)** 꺼지고 

   - → Antigravity가 이 폴더에서 코드 생성/수정/실행을 제대로 할 수 있어. 

- 이 버튼 누르고 나면 창이 사라지고, 

   - 왼쪽에 `n1_crawler` 폴더 구조 그대로 뜨고 

   - 오른쪽 Agent 패널에서 이제 아까 말한 **긴 프롬프트** 붙여넣는 단계로 넘어가면 된다. 

**==> picture [289 x 140] intentionally omitted <==**

확인해줘 

좋아 형, 지금 화면 완전 깔끔하게 잘 됐어. 👍 **폴더/파일 상태부터 체크해 줄게.** 

## **1. 지금 상태 체크** 

왼쪽 `n1_crawler` 안에: 

- ✅ 새로 만든 애들 

   - `build_exe.bat cleaner.py fetcher.py logger.py main.py patterns.py post.py` 

   - `requirements.txt` 

   - `saver.py` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

30/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
tagger.py
```

- 🟡 예전에 쓰던 애들(안 써도 되는 파일) 

`analyzer.py config.py crawler.py crawler_log.txt __pycache__` 폴더 

- 👉 이 네 개랑 `__pycache__` 는 **이번 v2.4 크롤러랑은 상관 없는 옛날 버전** 이라 헷갈리지 않으려면 **지워도 된다.** 

정리 추천 

- “ `analyzer.py` , `config.py` , `crawler.py` , `crawler_log.txt` , `__pycache__` → 삭제 (마우스 오른쪽 → Delete)” 

그렇게 하면 딱 우리가 설계한 **10개 파일만 있는 폴더** 가 된다. 

## **2. 다음 단계: 라이브러리 설치 (requirements.txt)** 

이제 **실행 준비** 하자. 

Antigravity 안에서 바로 하면 편해. 

**1.** 위 메뉴에서 `Terminal` **→** `New Terminal` 클릭. 

**2.** 아래쪽에 검은(또는 어두운) 터미널 창이 뜨고, 프롬프트가 대략 이렇게 보이면 OK: 

text 코드 복사 

```
C:\Users\psm23\Desktop\n1_crawler>
```

**3.** 여기서 아래 명령 입력하고 엔터: 

bash 코드 복사 `pip install -r requirements.txt` 

**4.** 설치가 쭉 지나가고 에러만 안 뜨면 끝. 

## **3. 파이썬으로 먼저 테스트 실행** 

exe 만들기 전에 **main.py가 잘 도는지** 한 번 확인하자. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

31/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

같은 터미널에서: 

bash 코드 복사 

```
python main.py
```

화면에 text 코드 복사 `============================= 엔오원 니즈 수집 크롤러 v2 ============================= [1] 사이트 이름 입력:` 

이런 식으로 나오면 성공. 테스트용으로 

사이트 이름: `test_site` URL: 테스트용 게시판 주소 1개만 넣고 

더 이상 없으면 빈 줄에서 엔터. 

끝나고 나서 

`output` 폴더 자동 생성 `output/test_site.csv` 생기면 **정상 작동** . 

## **4. exe 만들기 (build_exe.bat 실행)** 

테스트까지 OK면 이제 exe로 빌드. 

**1.** 터미널이 아직 `n1_crawler` 폴더 안이면, 그대로: 

bash 코드 복사 `build_exe.bat` bash 코드 복사 `.\build_exe.bat` 

또는 

**2.** 안에 `pyinstaller --onefile --name n1_crawler main.py` 가 실행되고, 끝나면 폴더 구조에: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

32/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

`build` 폴더 

`dist` 폴더 안에 `n1_crawler.exe` 생김. 

## **5. 최종 사용 방법** 

**1.** 윈도우 탐색기에서 `n1_crawler\dist\n1_crawler.exe` 더블클릭. 

**2.** 콘솔창에서: 

[1] 사이트 이름 입력 

[2] URL들 한 줄씩 입력 → 마지막은 빈 줄 엔터. 

**3.** 끝나고 

`output/사이트이름.csv` 만들어지면 성공. 

## 지금 화면 기준으로는: 

구조 ✅ 

코드도 ( `cleaner.py` 탭 내용 보니까) 우리가 말한 규칙 그대로라 ✅ 

이제 **터미널 열어서** `pip install -r requirements.txt` **→** `python main.py` **순서** 로 한 번 돌려보고, 