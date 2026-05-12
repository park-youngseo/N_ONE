프로그램[파이썬/코딩] - [Voice] Script 생성<br>**----- End of picture text -----**<br>


enable-libbs2b --enable-libflite --enable-libmysofa -- enable-librubberband --enable-libsoxr --enablechromaprint --enable-whisper libavutil      60.  8.100 / 60.  8.100 libavcodec     62. 11.100 / 62. 11.100 libavformat    62.  3.100 / 62.  3.100 libavdevice    62.  1.100 / 62.  1.100 libavfilter    11.  4.100 / 11.  4.100 libswscale      9.  1.100 /  9.  1.100 libswresample   6.  1.100 /  6.  1.100 

Exiting with exit code 0 

C:\Users\psm23> 

형, 잘했어! **FFmpeg 설치 완료** 야. 이제 바로 “음성→텍스트” 파이프라인 돌려보자. (복사→붙여넣기만 하면 돼) 

## **1) 파이썬 준비(한 줄씩)** 

nginx 코드 복사 

```
pip install openai webrtcvad pandas
```

(화자 구분까지 쓸 거면) 

nginx 코드 복사 `pip install pyannote.audio torch` 그리고 키 넣기(있다면): nginx 코드 복사 `setx OPENAI_API_KEY "여기_본인_API키" setx HUGGINGFACE_TOKEN "여기_본인_토큰"   ← 화자 구분 쓸 때만` 

그리고 키 넣기(있다면): 

※ 키 넣은 뒤엔 **창을 닫았다 열어야** 인식돼. 

## **2) 파이프라인 파일 받기** 

stt_pipeline.py 다운로드 

다운로드 폴더에 저장해 둔 뒤, 명령창에서: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

55/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

코드 복사 

bash 

 

```
cd %USERPROFILE%\Downloads
```

## **3) 바로 실행(예시: 회의.mp3)** 

기본(보통 전처리 + 60초 분할): 

**==> picture [481 x 322] intentionally omitted <==**

**----- Start of picture text -----**<br>
nginx 코드 복사<br>python stt_pipeline.py 회의.mp3<br>소리 작음/멀어짐 보정  강모드  + 음악·소음 제거( 대화만 ):<br>css 코드 복사<br>python stt_pipeline.py 회의.mp3 --strong --vad<br>참여자 구분 (사람별 태그 붙이기):<br>css 코드 복사<br>python stt_pipeline.py 회의.mp3 --diarize<br>(필요하면 같이)<br>css 코드 복사<br>**----- End of picture text -----**<br>


```
python stt_pipeline.py회의.mp3--diarize--vad--strong
```

## **4) 결과물(같은 폴더에 생김)** 

`회의_transcript.txt` ← 줄글 

`회의.srt` ← 자막(시간표 포함) 

- `회의_segments.csv` ← 시작/끝시간·(화자)·문장 표 

## **5) 안 될 때 딱 3가지만 확인** 

**1.** `ffmpeg -version` → 숫자 나오면 OK (이미 됨) 

**2.** `python` 대신 `py` 가 필요한 PC도 있어: 

css 코드 복사 

```
py -m pip install openai webrtcvad pandas
py stt_pipeline.py회의.mp3--strong--vad
```

## **3.** 키 넣었으면 **창 닫고 다시 열기** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

56/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

원하면 “메뉴형(번호 고르기)” 초간단 버전으로 바꿔줄게. 지금 이대로 먼저 한 번 돌려보 고, 어디서 막혔는지만 알려줘! 

**==> picture [489 x 38] intentionally omitted <==**

C:\Users\psm23>pip install openai webrtcvad pandas Collecting openai 

Downloading openai-2.6.1-py3-none-any.whl.metadata (29 kB) 

Collecting webrtcvad 

Downloading webrtcvad-2.0.10.tar.gz (66 kB) 

---------------------------------------- 66.2/66.2 kB 1.2 

MB/s eta 0:00:00 

Installing build dependencies ... done Getting requirements to build wheel ... done Preparing metadata (pyproject.toml) ... done Requirement already satisfied: pandas in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (2.3.3) Collecting anyio<5,>=3.5.0 (from openai) 

Downloading anyio-4.11.0-py3-none-any.whl.metadata (4.1 kB) 

Collecting distro<2,>=1.7.0 (from openai) Downloading distro-1.9.0-py3-none-any.whl.metadata (6.8 kB) 

Collecting httpx<1,>=0.23.0 (from openai) Downloading httpx-0.28.1-py3-none-any.whl.metadata (7.1 kB) 

Collecting jiter<1,>=0.10.0 (from openai) Downloading jiter-0.11.1-cp311-cp311win_amd64.whl.metadata (5.3 kB) Collecting pydantic<3,>=1.9.0 (from openai) Downloading pydantic-2.12.3-py3-noneany.whl.metadata (87 kB) ---------------------------------------- 87.7/87.7 kB 980.0 kB/s eta 0:00:00 Collecting sniffio (from openai) Downloading sniffio-1.3.1-py3-none-any.whl.metadata (3.9 kB) Collecting tqdm>4 (from openai) Downloading tqdm-4.67.1-py3-none-any.whl.metadata 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

57/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

(57 kB) 

---------------------------------------- 57.7/57.7 kB 

608.1 kB/s eta 0:00:00 Requirement already satisfied: typingextensions<5,>=4.11 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from openai) (4.15.0) Requirement already satisfied: numpy>=1.23.2 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from pandas) (2.3.4) Requirement already satisfied: python-dateutil>=2.8.2 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from pandas) (2.9.0.post0) Requirement already satisfied: pytz>=2020.1 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from pandas) (2025.2) Requirement already satisfied: tzdata>=2022.7 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from pandas) (2025.2) Requirement already satisfied: idna>=2.8 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from anyio<5,>=3.5.0>openai) (3.11) 

Requirement already satisfied: certifi in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpx<1,>=0.23.0>openai) (2025.10.5) Collecting httpcore==1.* (from httpx<1,>=0.23.0>openai) Downloading httpcore-1.0.9-py3-noneany.whl.metadata (21 kB) Collecting h11>=0.16 (from httpcore==1.*>httpx<1,>=0.23.0->openai) Downloading h11-0.16.0-py3-none-any.whl.metadata (8.3 kB) Collecting annotated-types>=0.6.0 (from pydantic<3,>=1.9.0->openai) Downloading annotated_types-0.7.0-py3-noneany.whl.metadata (15 kB) Collecting pydantic-core==2.41.4 (from pydantic<3,>=1.9.0->openai) Downloading pydantic_core-2.41.4-cp311-cp311win_amd64.whl.metadata (7.4 kB) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

58/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

Collecting typing-inspection>=0.4.2 (from pydantic<3,>=1.9.0->openai) Downloading typing_inspection-0.4.2-py3-noneany.whl.metadata (2.6 kB) Requirement already satisfied: six>=1.5 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from pythondateutil>=2.8.2->pandas) (1.17.0) Requirement already satisfied: colorama in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from tqdm>4->openai) (0.4.6) 

Downloading openai-2.6.1-py3-none-any.whl (1.0 MB) ---------------------------------------- 1.0/1.0 MB 707.4 

kB/s eta 0:00:00 

Downloading anyio-4.11.0-py3-none-any.whl (109 kB) ---------------------------------------- 109.1/109.1 kB 

790.9 kB/s eta 0:00:00 

Downloading distro-1.9.0-py3-none-any.whl (20 kB) Downloading httpx-0.28.1-py3-none-any.whl (73 kB) ---------------------------------------- 73.5/73.5 kB 

673.2 kB/s eta 0:00:00 

Downloading httpcore-1.0.9-py3-none-any.whl (78 kB) ---------------------------------------- 78.8/78.8 kB 

729.2 kB/s eta 0:00:00 

Downloading jiter-0.11.1-cp311-cp311-win_amd64.whl (207 kB) 

---------------------------------------- 207.6/207.6 kB 

789.4 kB/s eta 0:00:00 

Downloading pydantic-2.12.3-py3-none-any.whl (462 kB) ---------------------------------------- 462.4/462.4 kB 

117.2 kB/s eta 0:00:00 Downloading pydantic_core-2.41.4-cp311-cp311win_amd64.whl (2.0 MB) 

---------------------------------------- 2.0/2.0 MB 715.4 kB/s eta 0:00:00 

Downloading sniffio-1.3.1-py3-none-any.whl (10 kB) Downloading tqdm-4.67.1-py3-none-any.whl (78 kB) ---------------------------------------- 78.5/78.5 kB 

726.6 kB/s eta 0:00:00 

Downloading annotated_types-0.7.0-py3-none-any.whl (13 kB) 

Downloading typing_inspection-0.4.2-py3-none-any.whl 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

59/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

(14 kB) Downloading h11-0.16.0-py3-none-any.whl (37 kB) Building wheels for collected packages: webrtcvad Building wheel for webrtcvad (pyproject.toml) ... error error: subprocess-exited-with-error 

× Building wheel for webrtcvad (pyproject.toml) did not run successfully. 

│ exit code: 1 

╰─> [23 lines of output] 

C:\Users\psm23\AppData\Local\Temp\pipbuild-env-a8wbjclc\overlay\Lib\sitepackages\setuptools\_distutils\dist.py:289: UserWarning: Unknown distribution option: 'test_suite' warnings.warn(msg) 

C:\Users\psm23\AppData\Local\Temp\pipbuild-env-a8wbjclc\overlay\Lib\sitepackages\setuptools\dist.py:759: 

SetuptoolsDeprecationWarning: License classifiers are deprecated. 

!! 

************************************************************* 

******************* 

Please consider removing the following classifiers in favor of a SPDX license expression: 

License :: OSI Approved :: MIT License 

See 

https://packaging.python.org/en/latest/guides/writingpyproject-toml/#license for details. 

************************************************************* 

******************* 

!! 

self._finalize_license_expression() running bdist_wheel running build running build_py creating build\lib.win-amd64-cpython-311 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

60/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

copying webrtcvad.py -> build\lib.win-amd64cpython-311 running build_ext building '_webrtcvad' extension error: Microsoft Visual C++ 14.0 or greater is required. Get it with "Microsoft C++ Build Tools": https://visualstudio.microsoft.com/visual-cpp-build-tools/ [end of output] 

note: This error originates from a subprocess, and is likely not a problem with pip. ERROR: Failed building wheel for webrtcvad Failed to build webrtcvad ERROR: Could not build wheels for webrtcvad, which is required to install pyproject.toml-based projects 

[notice] A new release of pip is available: 24.0 -> 25.3 [notice] To update, run: python.exe -m pip install -- upgrade pip 

38s 동안 생각함 

형, 잘 봤어. **오류 원인** 은 `webrtcvad` 가 **C++ 빌드 도구** 가 없어서 설치를 못 한 거야. 초등학생 버전으로, **가장 쉬운 해결 2가지** 를 줄게. (하나는 1분 컷) 

## **방법 ① (추천·1분 컷) ― 준비된 버전 쓰기** 

빌드가 필요 없는 **완제품** 을 깔면 끝이야. 

**1.** 명령창에서 한 줄씩: 

nginx 코드 복사 `pip uninstall -y webrtcvad pip install webrtcvad-wheels` 끝나면 우리 파이프라인 그대로 실행: css 코드 복사 

**2.** 끝나면 우리 파이프라인 그대로 실행: 

```
python stt_pipeline.py회의.mp3--strong--vad
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

61/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

`webrtcvad-wheels` 는 같은 이름의 모듈( `import webrtcvad` )을 똑같이 제공해. 그래서 코드 수정 없이 돌아가. 

## **방법 ② (정석) ― C++ 빌드 도구 설치 후 원본 설치** 

조금 오래 걸리지만, 한 번만 하면 돼. 
