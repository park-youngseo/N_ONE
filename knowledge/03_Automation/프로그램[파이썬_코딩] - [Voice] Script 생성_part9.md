
25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

방법 1(화면으로): 

**시작 → 설정 → 시스템 → 정보 → 고급 시스템 설정 → 환경 변수** 

**사용자 변수** 의 **Path** 선택 → **새로 만들기** → `C:\ffmpeg\bin` 입력 → 확인 연달아 누르기 

방법 2(명령어, 일반 권한): 

**시작 → cmd** 열고 아래 입력: 

nginx 코드 복사 

```
setx PATH "%PATH%;C:\ffmpeg\bin"
```

창 **닫고** 새로 cmd 열기(중요!) 

## **4.** 확인: 

nginx 

코드 복사 

```
ffmpeg -version
```

버전 정보가 뜨면 성공! 

## **맥 / 리눅스(짧게)** 

## **macOS (Homebrew)** : 

nginx 코드 복사 `brew install ffmpeg ffmpeg -version` 

## **Ubuntu/Debian** : 

pgsql 

코드 복사 

```
sudo apt update
sudo apt install ffmpeg
ffmpeg -version
```

## **자주 생기는 문제 (5초 해결)** 

`ffmpeg` 가 **명령어로 인식 안 됨** → Path 등록 후 **새 창** (cmd/PowerShell)을 열어 다시 시도. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 49/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## `권한` 문제 → PowerShell을 **관리자** 로 다시 실행. 

설치는 됐는데 우리 스크립트가 못 찾음 → `ffmpeg -version` 이 되는 창에서 스크립 트를 실행. 

여기까지 되면, 아까 준 전처리/분할/대화추출/전사 스크립트가 전부 잘 돌아가. 설치 중에 막히는 화면이 있으면 딱 그 부분을 말해줘—바로 거기부터 잡아줄게. 

**==> picture [489 x 38] intentionally omitted <==**

진짜 너 죽을래? 내가 초등학생도 설치 할수 있도록 알려 달라고 했지? 

25s 동안 생각함 

형, 미안해. 이번엔 **진짜 초등학생도 따라 하는** 방식으로 알려줄게. (윈도우 컴퓨터 기준) 

## **FFmpeg 설치 ─ 가장 쉬운 방법(그림 없이도 되는 완전 단 계)** 

## **방법 1) 글자 한 줄로 끝내기** 

**1.** 화면 왼쪽 아래 **시작** 버튼 누르기 → 검색칸에 `PowerShell` 쓰기. 

**2.** 파란 아이콘인 **Windows PowerShell** 위에서 **오른쪽 클릭** → **관리자 권한으로 실행** 누르기. 

**3.** 파란 창이 뜨면, 아래 글자를 **그대로 붙여넣고 엔터** : 

bash 코드 복사 

```
winget install --id=FFmpeg.FFmpeg -e
```

**4.** 설치가 끝나면 파란 창 닫기. 

**5.** 다시 시작 버튼 → `cmd` 입력 → **명령 프롬프트** 열기 → 아래 글자 치고 엔터: 

nginx 코드 복사 

```
ffmpeg -version
```

숫자(버전 정보)가 보이면 **성공** 이야! 

만약 “winget을 찾을 수 없음”이라고 나오면 **방법 2** 로 하면 돼. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

50/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **방법 2) 압축파일 받아서 설치(조금 길지만 확실)** 

**1. 다운로드 폴더** 에 파일을 받는다. (검색창에 `FFmpeg download` 라고 치고, **Windows 용 ZIP(압축파일)**을 받아.) 

**2. C:드라이브** 에 **새 폴더** 만들기 → 이름을 ** `ffmpeg` **로 적기 → 결과: `C:\ffmpeg` 

**3.** 방금 받은 **압축파일을 더블클릭** → 안에 있는 폴더를 **통째로** `C:\ffmpeg` **에 끌어넣기** 최종 모양 예시: `C:\ffmpeg\bin\ffmpeg.exe` 가 보이면 OK 

**4.** 이제 **경로(Path) 등록** 하기 

   **1. 시작 → 설정 → 시스템 → 정보 → 고급 시스템 설정** 누르기 

   **2. 환경 변수** 버튼 누르기 

   **3.** 위쪽 **사용자 변수** 에서 **Path** 를 두 번 클릭 

   **4. 새로 만들기** 버튼 누르고 `C:\ffmpeg\bin` 적기 → **확인** 계속 눌러 창 닫기 

**5. 새로** 명령 프롬프트(cmd) 열기 → 아래 치고 엔터: 

nginx 코드 복사 

```
ffmpeg -version
```

숫자(버전 정보)가 보이면 **성공** ! 

## **안 될 때 빠른 해결(딱 3개)** 

- 창을 **다시 열기** : Path를 추가했으면 **새 창** 을 열어야 인식돼. 

- 철자 확인: `C:\ffmpeg\bin` 오타 없는지 보기. 

- 관리자 권한: **PowerShell을 관리자** 로 열고 방법 1 다시 실행. 

여기까지 되면, 우리가 만든 **전처리/분할/대화만 추출/자막 변환** 스크립트가 다 돌아가. 설치하다가 **딱 어디에서 멈췄는지** 한 줄만 말해줘. 거기서 바로 이어서 잡아줄게. 

**==> picture [489 x 38] intentionally omitted <==**

## Windows PowerShell 

Copyright (C) Microsoft Corporation. All rights reserved. 

새로운 기능 및 개선 사항에 대 한 최신 PowerShell을 설 치 하세요! https://aka.ms/PSWindows 

PS C:\WINDOWS\System32> winget install -- 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

51/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

id=FFmpeg.FFmpeg -e 

>> 

입력 조건과 일치하는 패키지를 찾을 수 없습니다. PS C:\WINDOWS\System32> 

40s 동안 생각함 

오케이 형, 지금 오류는 **winget이 FFmpeg 패키지를 못 찾는** 상황이야. 아래 순서 그대로 해보자. (복사→붙여넣기만 하면 됨) 

## **1단계: winget이 FFmpeg 찾게 만들기** 

PowerShell(관리자)에서 한 줄씩: 

sql 코드 복사 `winget` _`--version`_ `winget source update winget search ffmpeg` 

마지막 줄에 목록이 보이면, 설치: 

bash 코드 복사 `winget install -e --id Gyan.FFmpeg` 

안 되면(여전히 “못 찾음” 나오면) 이걸로 시도: 

nginx 코드 복사 `winget install ffmpeg` 

## **2단계: 설치 확인** 

새로 **명령프롬프트(cmd)** 열고: 

nginx 코드 복사 `ffmpeg -version` 

버전 숫자 나오면 성공! 

## **그래도 안 되면(회사 PC, winget 막힘 등) — 쉬운 우회 2가지** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

52/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **방법 A) Scoop로 설치(매우 간단)** 

PowerShell(일반 권한)에서: 

powershell 코드 복사 `Set-ExecutionPolicy RemoteSigned -Scope CurrentUser iwr -useb get.scoop.sh | iex scoop install ffmpeg` 

확인: 

nginx 코드 복사 `ffmpeg -version` 

## **방법 B) 수동 설치(압축파일)** 

**1.** `C:\ffmpeg` 폴더 만들기 

**2.** 이 사이트에서 Windows용 ZIP 받기 → 풀기 → `C:\ffmpeg\bin\ffmpeg.exe` 가 보이게 정리 

(빌드 제공처: gyan.dev) gyan.dev 

**3. 환경변수 Path** 에 `C:\ffmpeg\bin` 추가 

PowerShell에서: 

nginx 코드 복사 

```
setx PATH "$env:PATH;C:\ffmpeg\bin"
```

## **4.** 새 cmd 열고: 

nginx 코드 복사 

```
ffmpeg -version
```

## **참고(왜 “Gyan.FFmpeg”인가?)** 

winget에서 FFmpeg는 보통 **Gyan.FFmpeg** ID로 배포돼. 설치 명령 예시는 공식 패키지 정보 페이지에도 있어. winget.run 

가끔 winget 쪽 이슈로 인식이 꼬일 때가 있어서( `source update` 로 해결되는 경우 많음). GitHub +1 

## 막히는 **화면/문구** 를 그대로 붙여주면, 그 지점부터 딱 잡아 줄게. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

53/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

**출처** 

**==> picture [12 x 12] intentionally omitted <==**

**==> picture [11 x 12] intentionally omitted <==**

Microsoft Windows [Version 10.0.26100.6899] (c) Microsoft Corporation. All rights reserved. 

C:\Users\psm23>ffmpeg -version 

ffmpeg version 8.0-full_build-www.gyan.dev Copyright (c) 2000-2025 the FFmpeg developers built with gcc 15.2.0 (Rev8, Built by MSYS2 project) configuration: --enable-gpl --enable-version3 --enablestatic --disable-w32threads --disable-autodetect -- enable-fontconfig --enable-iconv --enable-gnutls -- enable-lcms2 --enable-libxml2 --enable-gmp --enablebzlib --enable-lzma --enable-libsnappy --enable-zlib -- enable-librist --enable-libsrt --enable-libssh --enablelibzmq --enable-avisynth --enable-libbluray --enablelibcaca --enable-libdvdnav --enable-libdvdread -- enable-sdl2 --enable-libaribb24 --enable-libaribcaption - -enable-libdav1d --enable-libdavs2 --enable-libopenjpeg --enable-libquirc --enable-libuavs3d --enable-libxevd -- enable-libzvbi --enable-liboapv --enable-libqrencode -- enable-librav1e --enable-libsvtav1 --enable-libvvenc -- enable-libwebp --enable-libx264 --enable-libx265 -- enable-libxavs2 --enable-libxeve --enable-libxvid -- enable-libaom --enable-libjxl --enable-libvpx --enablemediafoundation --enable-libass --enable-frei0r -- enable-libfreetype --enable-libfribidi --enablelibharfbuzz --enable-liblensfun --enable-libvidstab -- enable-libvmaf --enable-libzimg --enable-amf --enablecuda-llvm --enable-cuvid --enable-dxva2 --enabled3d11va --enable-d3d12va --enable-ffnvcodec --enablelibvpl --enable-nvdec --enable-nvenc --enable-vaapi -- enable-libshaderc --enable-vulkan --enable-libplacebo -- enable-opencl --enable-libcdio --enable-openal -- enable-libgme --enable-libmodplug --enablelibopenmpt --enable-libopencore-amrwb --enablelibmp3lame --enable-libshine --enable-libtheora -- enable-libtwolame --enable-libvo-amrwbenc --enablelibcodec2 --enable-libilbc --enable-libgsm --enableliblc3 --enable-libopencore-amrnb --enable-libopus -- enable-libspeex --enable-libvorbis --enable-ladspa -- 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

54/175 

25. 11. 30. 오후 1:03 

 

**==> picture [145 x 9] intentionally omitted <==**

**----- Start of picture text -----**<br>