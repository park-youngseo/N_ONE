형, 아주 쉬운 말로 “처음부터” 설명할게. 

지금 올린 **파이썬_01.txt** 안에는 **완전 초보** 도 바로 따라 할 수 있는 4가지가 있어: 

**1. 영상 재생 웹페이지** , 2) **이미지 보기 웹페이지** , 3) **음성 → 글자(자막) 만들기** , 4) **영어 → 한국어 번역** . 

이 4가지는 “수업·홍보·정리”에 바로 써먹을 수 있어. 

## **0. 먼저 준비(5분)** 

- **파이썬** 설치: 공식 사이트에서 설치(다음, 다음만 눌러도 됨). 

- **명령창** 열기: 윈도우 키 누르고 `cmd` 입력 → 엔터. 

- 아래 한 줄씩 복사해서 엔터: 

nginx 코드 복사 

```
pip install streamlit pillow
```

```
pip install openai deepl googletrans==4.0.0-rc1
```

- (※ 번역·자막에 **유료 API키** 가 필요할 수 있어. 키가 없으면 1)영상 2)이미지부터 시작 해도 충분해.) 

## **1. 영상 재생 웹페이지 (가장 쉬움)** 

**뭐야?** : 유튜브 링크나 내 PC의 동영상을 **브라우저에서 바로 재생** 하는 아주 간단한 사 이트. 

**어디에 써?** : 수업 자료 보여주기, 홍보 영상 데모. 

## **따라하기** 

**1.** 바탕화면에 `video_app.py` 파일을 하나 만든다. 

**2.** 파일 안에 아래를 그대로 붙여넣기: 

python 코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

7/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
import streamlit as st
```

```
st.title("영상재생기")
tab1, tab2 = st.tabs(["유튜브링크", "내컴퓨터파일"])
with tab1:
  url = st.text_input("유튜브링크를붙여넣으세요")
if url:
    st.video(url)
with tab2:
  file = st.file_uploader("동영상파일선택(mp4 등)", type=["mp4","mov","m4
if file: st.video(file)
```

  명령창에서: arduino 코드 복사 

**3.** 명령창에서: 

```
cd %USERPROFILE%\Desktop
streamlit run video_app.py
```

## → 브라우저가 열리면 **링크 붙여넣기** 또는 **파일 올리기** 만 하면 끝. 

## **2. 이미지 보기 웹페이지** 

**뭐야?** : 내 컴퓨터 이미지와 인터넷 이미지(URL)를 **한 화면에서 크게 보기** . 

**어디에 써?** : 포스터/썸네일 비교, 과제 예시 보여주기. 

**따라하기** 

**1.** `image_app.py` 만들기 → 아래 코드 붙여넣기: 

python 코드 복사 

```
import streamlit as st
from PIL import Image
import requests
from io import BytesIO
st.title("이미지보기")
col1, col2 = st.columns(2)
with col1:
  st.subheader("내컴퓨터")
  f = st.file_uploader("이미지선택", type=["png","jpg","jpeg"])
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

8/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
if f:
    img = Image.open(f)
    st.image(img, use_container_width=True)
with col2:
  st.subheader("인터넷이미지")
  url = st.text_input("이미지주소(URL) 붙여넣기")
if url:
try:
      data = requests.get(url, timeout=5).content
      st.image(Image.open(BytesIO(data)), use_container_width=True)
except: st.error("이미지주소가이상해요.")
```

**2.** 실행: 

arduino 코드 복사 

```
cd %USERPROFILE%\Desktop
streamlit run image_app.py
```

## **3. 음성 → 글자(자막) 만들기** 

- **뭐야?** : mp3 같은 **녹음 파일을 글자로 바꿔** 준다. (회의록, 강의 요약에 최고) **어디에 써?** : 수업/회의 녹음 → 텍스트 정리. 

- **준비** : OpenAI(또는 다른 서비스) **API 키** 가 필요. 키가 없으면 이 단계는 잠시 건너뛰 기. 

## **따라하기(가장 단순 버전)** 

**1.** `transcribe.py` 만들고 아래 붙여넣기: 

python 코드 복사 `import os from pathlib import Path from openai import OpenAI` 

```
# 1) 환경변수에 API키넣기: setx OPENAI_API_KEY "발급받은_키"
```

```
client = OpenAI()
audio_path = "sample.mp3"# 같은폴더에 mp3 파일두기
withopen(audio_path, "rb") as f:
    result = client.audio.transcriptions.create(
        model="whisper-1",
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

9/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
        file=f
```

```
    )
```

```
text = result.text
Path("result.txt").write_text(text, encoding="utf-8")
print("완료! result.txt 를확인하세요.")
```

**2.** 명령창: 

bash 코드 복사 

```
setx OPENAI_API_KEY "여기에_본인_API키"
cd %USERPROFILE%\Desktop
python transcribe.py
```

## **3. result.txt** 생기면 성공. 

## **4. 영어 → 한국어 번역** 

**뭐야?** : 긴 영어 글도 **한 번에 한국어** 로 바꿔 준다. 

**어디에 써?** : 해외 자료 번역, 유튜브 자막 번역. 

**준비** : DeepL이나 OpenAI 같은 **API 키** 가 있으면 좋아. 

**따라하기(DeepL 예시)** 

**1.** 키 받으면: 

nginx 코드 복사 `setx DEEPL_API_KEY "본인_딥엘_키"` **2.** `translate.py` 만들고 붙여넣기: 

python 코드 복사 

```
import os
import deepl
```

```
auth = os.environ.get("DEEPL_API_KEY")
translator = deepl.Translator(auth)
```

```
english = "Hello! This is a simple test for translation."
result = translator.translate_text(english, target_lang="KO")
print(result.text)
```

## **3.** 실행: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 10/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

코드 복사  

bash 

```
cd %USERPROFILE%\Desktop
```

```
python translate.py
```

## **어디에 적용하면 좋아?** 

- **수업/학원** : 영상·이미지 **자료실 웹페이지** 로 바로 보여주기, 녹음 파일 **텍스트 정리** . **홍보/마케팅** : 제품/작품 영상 데모, 해외 글 **번역** 해 참고. 

- **팀 운영** : 회의 녹취 정리, 공지 다국어 번역. 

## **오늘 바로 해볼 “10분 코스”** 

**1. 영상 재생** : `video_app.py` 만들고 `streamlit run` 으로 띄워서 유튜브 링크 재생. 

**2. 이미지 보기** : `image_app.py` 띄워서 사진 파일/URL 둘 다 보여보기. 

**3.** (API키 있으면) **음성 → 글자** 돌려서 `result.txt` 확인. 

**4.** (키 있으면) **번역** 코드로 한 문장 번역 출력 확인. 

- 형, 여기까지 그대로 따라 하면 **화면에 바로 보이는 웹앱 2개** 랑 **텍스트 자동화 2개** 를 손에 쥐게 돼. 

- “다음”이라고 하면, 위 4개를 **한 페이지 메뉴** 로 묶은 **올인원 대시보드** 까지 만들어줄게. 

**==> picture [489 x 39] intentionally omitted <==**

1-4까지의 예시가 있다면 링크를 보여줘 

13s 동안 생각함 

형, 1~4번 각각 **바로 보기 좋은 예시/설명 링크** 야. (눌러서 열면 돼!) 

## **1. 영상 재생 웹앱 (Streamlit)** 

- `st.video()` 공식 문서: 유튜브 링크·로컬 파일 재생 방법. docs.streamlit.io `st.file_uploader()` 문서: 동영상 파일 올리기 위젯. docs.streamlit.io 

## **2. 이미지 보기 웹앱 (Streamlit + Pillow)** 

- `st.image()` 공식 문서: 이미지 화면에 띄우는 법. docs.streamlit.io 

- Pillow 튜토리얼: `Image.open()` 등 기본 사용. Pillow (PIL Fork) 

## **3. 음성 → 글자(자막) (Whisper)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

11/175 

25. 11. 30. 오후 1:03 

 

   - 프로그램[파이썬/코딩] - [Voice] Script 생성 

- OpenAI 공식 “Speech-to-Text” 가이드: 파이썬 예시 포함. OpenAI 플랫폼 
