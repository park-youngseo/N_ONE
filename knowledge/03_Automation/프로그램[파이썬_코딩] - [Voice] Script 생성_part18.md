
https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

94/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

**==> picture [166 x 28] intentionally omitted <==**

**----- Start of picture text -----**<br>
파일은 폴더안에 넣어놨어<br>**----- End of picture text -----**<br>


**==> picture [481 x 687] intentionally omitted <==**

**----- Start of picture text -----**<br>
굿! 그럼 지금부터  딱 4줄 만 따라 해 👍<br>(바탕화면이  OneDrive\Desktop\Voice script   라고 했지?)<br>1) 폴더로 들어가기<br>bash 코드 복사<br>cd "%USERPROFILE%\OneDrive\Desktop\Voice script"<br>2) 파일 확인(스크립트랑 음성 둘 다 보이면 OK)<br>bash 코드 복사<br>dir<br>목록에  voice_script.py 랑  m4a/mp3 파일 이 보여야 해.<br>3) (처음이라면) 키가 필요한지 확인<br>아직  OpenAI 키 를 안 넣었다면 먼저 넣고, 창을  닫았다 다시 열고  위 1)로 돌아와:<br>nginx 코드 복사<br>setx OPENAI_API_KEY "여기_본인_API키"<br>(화자 구분 쓸 거면)<br>nginx 코드 복사<br>setx HUGGINGFACE_TOKEN "여기_본인_토큰"<br>4) 실행! (소리 보정 + 대화만)<br>css 코드 복사<br>python voice_script.py "네_파일이름.m4a" --strong --vad<br>예:<br>css 코드 복사<br>**----- End of picture text -----**<br>


https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

95/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong--vad
```

## **끝나면 이 3개가 같은 폴더에 생겨** 

`…_transcript.txt` (텍스트) 

- `….srt` (자막) 

- `…_segments.csv` (시간·문장 표) 

## **자주 막히는 3가지 (5초 해결)** 

**1. 키 에러** → 3)에서 키 넣고 **창을 닫았다 다시 열기** 

**2. 파일 못 찾음** → 파일명에 공백/대괄호 있으면 **반드시 "따옴표"** 

**3. python 안 됨** → `py` 로 바꿔 실행: 

css 

코드 복사 

```
py voice_script.py "네_파일이름.m4a" --strong--vad
```

실행했는데 메시지가 뜨면, **마지막 10줄** 만 붙여줘. 거기서 바로 해결해줄게! 

**==> picture [489 x 39] intentionally omitted <==**

Microsoft Windows [Version 10.0.26100.6899] (c) Microsoft Corporation. All rights reserved. 

C:\Users\psm23>mkdir 

"%USERPROFILE%\OneDrive\Desktop\Voice script" 하위 디렉터리 또는 파일 

C:\Users\psm23\OneDrive\Desktop\Voice script이(가) 이미 있습니다. 

C:\Users\psm23>cd 

"%USERPROFILE%\OneDrive\Desktop\Voice script" 

C:\Users\psm23\OneDrive\Desktop\Voice script>python voice_script.py "D:\녹음\[20230927]제이 타운_매매관련01.m4a" --strong --vad 파일 없음: D:\녹음\[20230927]제이타운_매매관련 01.m4a 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 96/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

C:\Users\psm23\OneDrive\Desktop\Voice script>dir C 드라이브의 볼륨에는 이름이 없습니다. 

볼륨 일련 번호: CC77-8F78 

C:\Users\psm23\OneDrive\Desktop\Voice script 디 렉터리 

2025-10-28  오전 12:40    <DIR>          . 

2025-10-28  오전 12:21    <DIR>          .. 

2025-10-28  오전 12:37             8,464 voice_script.py 2024-09-30  오후 06:23        26,643,745 [20230927]제이 타운_매매관련01.m4a 

2개 파일          26,652,209 바이트 

2개 디렉터리  953,278,775,296 바이트 남음 

C:\Users\psm23\OneDrive\Desktop\Voice script> 

형, 지금 **파일은 현재 폴더에 있는데** 실행할 때 **D:\녹음\…** 으로 잘못 적어서 “파일 없 음”이 뜬 거야. 

그냥 **현재 폴더의 파일이름만** 써서 돌리면 돼! 

## **바로 실행** 

css 

코드 복사 

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong--vad
```

대괄호가 있으니 **반드시 따옴표** 포함! 

## **실행 후 생기는 파일(같은 폴더)** 

- `[20230927]제이타운_매매관련01_transcript.txt` 

- `[20230927]제이타운_매매관련01.srt` 

- `[20230927]제이타운_매매관련01_segments.csv` 

## **만약 에러가 나오면** 

## **키 없음** 에러면: 

nginx 

코드 복사 

```
setx OPENAI_API_KEY "여기_본인_API키"
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 97/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

창을 **닫고 다시 열고** , 위 명령 다시 실행. 

**모듈 에러** 면: 

nginx 코드 복사 `pip install openai webrtcvad-wheels pandas python` 이 안 먹히면 `py` 로: css 코드 복사 

```
py voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong--vad
```

돌려보고, 메시지 마지막 **10줄** 만 붙여줘. 그 지점부터 바로 해결해줄게! 

**2/2** 

**==> picture [289 x 82] intentionally omitted <==**

새로 토큰을 생성하려고했지만 막혀있는데? 

42s 동안 생각함 

형, 화면에 **“Please confirm your account email address before creating access tokens”** 라고 보이지? 

즉, **이메일 인증을 먼저 해야 토큰을 만들 수 있어.** 아래 순서 그대로 해봐. 

## **1) 이메일 인증부터 (막힌 원인 해결)** 

**1.** Hugging Face 로그인 → 오른쪽 위 **프로필 사진 → Settings → Emails** . 

**2. Primary email** 옆에 **Verify / Resend verification** 버튼 눌러서 메일 재전송. 

**3.** 받은 메일에서 **Verify** 버튼 클릭. (스팸함도 확인) 

**4.** 같은 화면(Access Tokens)으로 돌아오면 이제 **+ Create new token** 이 활성화돼. 

   - 회사 메일이 보안으로 막히면, 집/모바일 네트워크에서 시도하거나 개인 메일로 변경 해도 돼. 

## **2) 토큰 만들기 (화자 구분용)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

98/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

**1. + Create new token** 클릭 

**2.** 이름: 아무거나 (예: `pyannote-read` ) 

**3.** 권한/Scope: **Read** 선택 

**4.** 생성되면 긴 문자열을 복사해서, 윈도우 명령창에: 

nginx 코드 복사 `setx HUGGINGFACE_TOKEN "여기_복사한_토큰"` 

**5. 창을 닫고 다시 열기** (중요), 확인: 

bash 코드 복사 `echo %HUGGINGFACE_TOKEN%` 

## **3) 모델 접근 허용(중요)** 

화자 구분은 `pyannote/speaker-diarization` 모델을 씀. 

토큰만 있다고 끝이 아니고, **모델 페이지에서 “Accept”** (약관 동의)를 눌러 **접근 허용** 을 해 야 해. 

- Hugging Face에서 `pyannote/speaker-diarization` 검색 → 페이지 들어가서 **“Access request/Agree/Accept”** 버튼 누르기. 

- 완료 후 다시 `--diarize` 실행. 

## **4) 지금 급하면 이렇게 먼저 돌려** 

토큰·접근 허용 끝나기 전에도 **대화만 추출(VAD)**로 충분히 정확도가 올라가. 

css 코드 복사 

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong--vad
```

나중에 화자 구분이 필요하면 같은 파일에 `--diarize` 만 추가해서 다시 실행하면 돼. 

## **5) 자주 생기는 막힘 한 번에 해결** 

**토큰 칸이 여전히 비활성** → 이메일 **인증 메일** 을 반드시 열고 **Verify** 클릭했는지 재확 인. 

- **조직/학교 메일 차단** → 개인 메일 추가 → **Make primary** → 그 메일로 인증. 

- **토큰 입력했는데 인식 안 됨** → `setx` 후 **명령창 새로 열기** 필수. 

- **모델 권한 오류(403/401)** → pyannote 모델 페이지에서 **Accept/Agree** 눌렀는지 확 인. 

필요하면 화면 캡처 한 장 더 보내줘. 어디에서 막혔는지 딱 짚어서 바로 풀어줄게! 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

99/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

**==> picture [272 x 193] intentionally omitted <==**

**==> picture [49 x 28] intentionally omitted <==**

**----- Start of picture text -----**<br>
뭐야<br>**----- End of picture text -----**<br>


형, 저 화면은 **pyannote 화자구분 모델을 쓰려면 짧은 신청서를 써달라** 는 거야. 세 칸만 채우면 돼. 부담 없이 **최소 정보** 로 써도 승인돼. 

## **이렇게 써** 

## **Company/university** : 

개인이면 그냥 `Personal` 또는 `Individual` 

## **Website** : 

사이트 없으면 `https://example.com` (형 GitHub나 블로그 있으면 그 주소도 OK) 

## **I plan to use this model for…** : 

예시 복붙 👉 `Meeting/phone recordings → speaker diarization for personal transcripts (Korean/English).` 

## 그리고 아래 **Agree and access repository** 누르면 끝. 

## **그다음 단계(한 번만)** 

## **1. 토큰 만들기** (이미 했으면 건너뜀) 

Hugging Face → Settings → **Access Tokens** → **New token** → **Read** → 복사 

**2. 토큰 넣기** (창 닫고 다시 열기 필수) 

nginx 

**==> picture [68 x 18] intentionally omitted <==**

**----- Start of picture text -----**<br>
코드 복사<br>**----- End of picture text -----**<br>


```
setx HUGGINGFACE_TOKEN "여기_복사한_토큰"
```