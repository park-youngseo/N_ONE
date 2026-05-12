
마지막 **10줄** 만 붙여줘. 거기서 바로 잡아줄게! 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

128/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

설치했는데 몇번을 새꺄 버전 확인을 제대로해 짜증나게 똑같은 상황이 반복되잖아 

Microsoft Windows [Version 10.0.26100.6899] 

(c) Microsoft Corporation. All rights reserved. 

C:\Users\psm23>python voice_script.py "[20230927]제 이타운_매매관련01.m4a" --strong --segment 60 --engine local --modelsize small 

python: can't open file 

'C:\\Users\\psm23\\voice_script.py': [Errno 2] No such file or directory 

C:\Users\psm23>cd "%USERPROFILE%\OneDrive\Desktop\Voice script" 

C:\Users\psm23\OneDrive\Desktop\Voice script>pip install faster-whisper 

Requirement already satisfied: faster-whisper in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (1.2.0) 

Requirement already satisfied: ctranslate2<5,>=4.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (4.6.0) Requirement already satisfied: huggingface-hub>=0.13 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (1.0.0) Requirement already satisfied: tokenizers<1,>=0.13 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (0.22.1) Requirement already satisfied: onnxruntime<2,>=1.14 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (1.23.2) Requirement already satisfied: av>=11 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (16.0.1) Requirement already satisfied: tqdm in 

c:\users\psm23\appdata\local\programs\python\py 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

129/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

thon311\lib\site-packages (from faster-whisper) (4.67.1) Requirement already satisfied: setuptools in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from ctranslate2<5,>=4.0>faster-whisper) (65.5.0) Requirement already satisfied: numpy in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from ctranslate2<5,>=4.0>faster-whisper) (2.3.4) Requirement already satisfied: pyyaml<7,>=5.3 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from ctranslate2<5,>=4.0>faster-whisper) (6.0.3) Requirement already satisfied: filelock in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (3.20.0) Requirement already satisfied: fsspec>=2023.5.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (2025.9.0) Requirement already satisfied: httpx<1,>=0.23.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (0.28.1) Requirement already satisfied: packaging>=20.9 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (25.0) Requirement already satisfied: shellingham in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (1.5.4) Requirement already satisfied: typer-slim in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (0.20.0) Requirement already satisfied: typingextensions>=3.7.4.3 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (4.15.0) Requirement already satisfied: hf-xet<2.0.0,>=1.2.0 in 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 130/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (1.2.0) Requirement already satisfied: coloredlogs in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from onnxruntime<2,>=1.14->faster-whisper) (15.0.1) Requirement already satisfied: flatbuffers in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from onnxruntime<2,>=1.14->faster-whisper) (25.9.23) Requirement already satisfied: protobuf in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from onnxruntime<2,>=1.14->faster-whisper) (6.33.0) Requirement already satisfied: sympy in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from onnxruntime<2,>=1.14->faster-whisper) (1.14.0) Requirement already satisfied: colorama in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from tqdm->faster-whisper) (0.4.6) 

Requirement already satisfied: anyio in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpx<1,>=0.23.0>huggingface-hub>=0.13->faster-whisper) (4.11.0) Requirement already satisfied: certifi in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpx<1,>=0.23.0>huggingface-hub>=0.13->faster-whisper) (2025.10.5) Requirement already satisfied: httpcore==1.* in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpx<1,>=0.23.0>huggingface-hub>=0.13->faster-whisper) (1.0.9) Requirement already satisfied: idna in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpx<1,>=0.23.0>huggingface-hub>=0.13->faster-whisper) (3.11) Requirement already satisfied: h11>=0.16 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpcore==1.*>httpx<1,>=0.23.0->huggingface-hub>=0.13->faster- 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

131/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

whisper) (0.16.0) 

Requirement already satisfied: humanfriendly>=9.1 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from coloredlogs>onnxruntime<2,>=1.14->faster-whisper) (10.0) Requirement already satisfied: mpmath<1.4,>=1.1.0 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from sympy>onnxruntime<2,>=1.14->faster-whisper) (1.3.0) Requirement already satisfied: click>=8.0.0 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from typer-slim>huggingface-hub>=0.13->faster-whisper) (8.3.0) Requirement already satisfied: pyreadline3 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from humanfriendly>=9.1>coloredlogs->onnxruntime<2,>=1.14->faster-whisper) (3.5.4) 

Requirement already satisfied: sniffio>=1.1 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from anyio- 

>httpx<1,>=0.23.0->huggingface-hub>=0.13->fasterwhisper) (1.3.1) 

[notice] A new release of pip is available: 24.0 -> 25.3 [notice] To update, run: python.exe -m pip install -- upgrade pip 

C:\Users\psm23\OneDrive\Desktop\Voice script>pip install torch --index-url https://download.pytorch.org/whl/cpu Looking in indexes: https://download.pytorch.org/whl/cpu Requirement already satisfied: torch in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (2.9.0) 

Requirement already satisfied: filelock in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (3.20.0) Requirement already satisfied: typingextensions>=4.10.0 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (4.15.0) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 132/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

Requirement already satisfied: sympy>=1.13.3 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (1.14.0) Requirement already satisfied: networkx>=2.5.1 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (3.5) Requirement already satisfied: jinja2 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (3.1.6) Requirement already satisfied: fsspec>=0.8.5 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (2025.9.0) Requirement already satisfied: mpmath<1.4,>=1.1.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from sympy>=1.13.3>torch) (1.3.0) 

Requirement already satisfied: MarkupSafe>=2.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from jinja2->torch) (3.0.3) 

[notice] A new release of pip is available: 24.0 -> 25.3 [notice] To update, run: python.exe -m pip install -- upgrade pip 

C:\Users\psm23\OneDrive\Desktop\Voice script>dir C 드라이브의 볼륨에는 이름이 없습니다. 볼륨 일련 번호: CC77-8F78 

C:\Users\psm23\OneDrive\Desktop\Voice script 디 렉터리 

2025-10-28  오전 02:07    <DIR>          . 2025-10-28  오전 02:05    <DIR>          .. 

2025-10-28  오전 01:48               279 API_Token.txt 2025-10-28  오전 02:05             3,577 run_voice_all.bat 2025-10-28  오전 02:15             9,420 voice_script.py 2024-09-30  오후 06:23        26,643,745 [20230927]제이 타운_매매관련01.m4a 

2025-10-28  오전 02:10        52,183,676 [20230927]제이 타운_매매관련01_clean.wav 

5개 파일          78,840,697 바이트 2개 디렉터리  951,254,626,304 바이트 남음 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 133/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

C:\Users\psm23\OneDrive\Desktop\Voice 

script>python voice_script.py "[20230927]제이타운_매매 관련01.m4a" --strong --segment 60 --engine local -- modelsize small 

[CMD] ffmpeg -y -i [20230927]제이타운_매매관련01.m4a -ac 1 -ar 16000 -af 

highpass=f=100,lowpass=f=8000,acompressor=threshold =-22dB:ratio=3:attack=5:release=250,volume=6dB,dynau dnorm=f=200:g=25,alimiter=limit=0.9 [20230927]제이타 운_매매관련01_clean.wav 

[전처리] [20230927]제이타운_매매관련01.m4a → 

[20230927]제이타운_매매관련01_clean.wav (강모드) 로컬 변환용 faster-whisper 설치 필요: pip install fasterwhisper 

C:\Users\psm23\OneDrive\Desktop\Voice script>pip install faster-whisper 

Requirement already satisfied: faster-whisper in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (1.2.0) 

Requirement already satisfied: ctranslate2<5,>=4.0 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (4.6.0) Requirement already satisfied: huggingface-hub>=0.13 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (1.0.0) Requirement already satisfied: tokenizers<1,>=0.13 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (0.22.1) Requirement already satisfied: onnxruntime<2,>=1.14 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (1.23.2) Requirement already satisfied: av>=11 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (16.0.1) Requirement already satisfied: tqdm in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (4.67.1) Requirement already satisfied: setuptools in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from ctranslate2<5,>=4.0>faster-whisper) (65.5.0) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

134/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

Requirement already satisfied: numpy in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from ctranslate2<5,>=4.0>faster-whisper) (2.3.4) Requirement already satisfied: pyyaml<7,>=5.3 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from ctranslate2<5,>=4.0>faster-whisper) (6.0.3) Requirement already satisfied: filelock in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (3.20.0) Requirement already satisfied: fsspec>=2023.5.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (2025.9.0) Requirement already satisfied: httpx<1,>=0.23.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (0.28.1) Requirement already satisfied: packaging>=20.9 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (25.0) Requirement already satisfied: shellingham in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (1.5.4) Requirement already satisfied: typer-slim in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (0.20.0) Requirement already satisfied: typingextensions>=3.7.4.3 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (4.15.0) Requirement already satisfied: hf-xet<2.0.0,>=1.2.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from huggingfacehub>=0.13->faster-whisper) (1.2.0) Requirement already satisfied: coloredlogs in c:\users\psm23\appdata\local\programs\python\py 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

135/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

thon311\lib\site-packages (from onnxruntime<2,>=1.14->faster-whisper) (15.0.1) Requirement already satisfied: flatbuffers in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from onnxruntime<2,>=1.14->faster-whisper) (25.9.23) Requirement already satisfied: protobuf in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from onnxruntime<2,>=1.14->faster-whisper) (6.33.0) Requirement already satisfied: sympy in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from onnxruntime<2,>=1.14->faster-whisper) (1.14.0) Requirement already satisfied: colorama in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from tqdm->faster-whisper) (0.4.6) 

Requirement already satisfied: anyio in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpx<1,>=0.23.0>huggingface-hub>=0.13->faster-whisper) (4.11.0) Requirement already satisfied: certifi in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpx<1,>=0.23.0>huggingface-hub>=0.13->faster-whisper) (2025.10.5) Requirement already satisfied: httpcore==1.* in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpx<1,>=0.23.0>huggingface-hub>=0.13->faster-whisper) (1.0.9) Requirement already satisfied: idna in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpx<1,>=0.23.0>huggingface-hub>=0.13->faster-whisper) (3.11) Requirement already satisfied: h11>=0.16 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from httpcore==1.*>httpx<1,>=0.23.0->huggingface-hub>=0.13->fasterwhisper) (0.16.0) 

Requirement already satisfied: humanfriendly>=9.1 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from coloredlogs>onnxruntime<2,>=1.14->faster-whisper) (10.0) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

136/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

Requirement already satisfied: mpmath<1.4,>=1.1.0 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from sympy>onnxruntime<2,>=1.14->faster-whisper) (1.3.0) Requirement already satisfied: click>=8.0.0 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from typer-slim- 

>huggingface-hub>=0.13->faster-whisper) (8.3.0) Requirement already satisfied: pyreadline3 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from humanfriendly>=9.1>coloredlogs->onnxruntime<2,>=1.14->faster-whisper) (3.5.4) 

Requirement already satisfied: sniffio>=1.1 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from anyio- 

>httpx<1,>=0.23.0->huggingface-hub>=0.13->fasterwhisper) (1.3.1) 

[notice] A new release of pip is available: 24.0 -> 25.3 [notice] To update, run: python.exe -m pip install -- upgrade pip 

C:\Users\psm23\OneDrive\Desktop\Voice script>python voice_script.py "[20230927]제이타운_매매 관련01.m4a" --strong --segment 60 --engine local -- modelsize small 

[CMD] ffmpeg -y -i [20230927]제이타운_매매관련01.m4a -ac 1 -ar 16000 -af 

highpass=f=100,lowpass=f=8000,acompressor=threshold =-22dB:ratio=3:attack=5:release=250,volume=6dB,dynau dnorm=f=200:g=25,alimiter=limit=0.9 [20230927]제이타 운_매매관련01_clean.wav 

[전처리] [20230927]제이타운_매매관련01.m4a → 

[20230927]제이타운_매매관련01_clean.wav (강모드) 로컬 변환용 faster-whisper 설치 필요: pip install fasterwhisper 

C:\Users\psm23\OneDrive\Desktop\Voice script>pip install faster-whisper 

Requirement already satisfied: faster-whisper in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (1.2.0) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

137/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

Requirement already satisfied: ctranslate2<5,>=4.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (4.6.0) Requirement already satisfied: huggingface-hub>=0.13 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (1.0.0) Requirement already satisfied: tokenizers<1,>=0.13 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (0.22.1) Requirement already satisfied: onnxruntime<2,>=1.14 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (1.23.2) Requirement already satisfied: av>=11 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (16.0.1) Requirement already satisfied: tqdm in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from faster-whisper) (4.67.1) Requirement already satisfied: setuptools in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from ctranslate2<5,>=4.0>faster-whisper) (65.5.0) 

Requirement already satisfied: numpy in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from ctranslate2<5,>=4.0>faster-whisper) (2.3.4) 

Requirement already satisfied: pyyaml<7,>=5.3 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from ctranslate2<5,>=4.0>faster-whisper) (6.0.3) 

Requirement already satisfied: filelock in 