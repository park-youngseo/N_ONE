                f.write(c)
# === 성공했을때만원본자막/중간 TXT 삭제 ===
for p in vtts + srts + txts:
try:
if os.path.exists(p):
                    os.remove(p)
except Exception as e:
                log_write(dest, f"[경고] 삭제실패: {p} -> {e}")
print("[완료] 통합및분할을마쳤습니다.")
        sys.exit(0)
except Exception as e:
# 오류시: 어떤것도삭제하지않음
try:
            log_write(dest if'dest'inlocals() else os.getcwd(), f"[오류] {repr(
except Exception:
pass
print(f"[오류] {e}")
        sys.exit(2)
if __name__ == "__main__":
    main()
```

## **사용 요약** 

**1.** 같은 폴더에 **Scirpt.bat** , **Youtube.py** 저장 (옵션: `cookies.txt` 함께). 

**2.** `Scirpt.bat` 실행 → 폴더명 입력 → URL들 붙여넣기 → 빈 줄 Enter. 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

32/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

**3.** 결과: `폴더이름_TXT.txt` (통합본) + `폴더이름_01.txt` …(4만자 분할). 

성공 시 원본 `.vtt/.srt/.txt` 는 삭제, 실패 시 파일 보존 + `Youtube_log.txt` 기록. 더 원하는 필터(이모지 제거, 말머리/말미 패턴 삭제 등) 있으면 말해줘. 

**==> picture [489 x 38] intentionally omitted <==**

두개의 파일이 제목이 똑같잖아 

28s 동안 생각함 

맞아. 저 두 파일은 `...ko.vtt` / `...ko-orig.vtt` 라서 **제목은 같지만 파일명 꼬리표** 가 달 라요. 

내가 준 스크립트에서 제목 키를 만들 때 `ko-orig` 까지 깎아내리지 못해서 서로 다른 제 목으로 인식한 거야. 아래처럼 **정규식 한 줄만 보강** 하면 두 파일이 **같은 제목으로 묶여** 하 나로 통합돼. 

## **패치 — Youtube.py만 수정** 

`extract_title_key` 함수의 언어코드 제거 부분을 이걸로 교체하세요. 

python 코드 복사 

```
# 언어/변형꼬리표제거 (예: .ko, .ko-KR, .ko-orig, .en, .zh-Hans 등)
#  - 확장자(.vtt/.srt/.txt)는위에서이미제거됨(os.path.splitext)
#  - 여기서는남아있는 ".ko-..." 같은꼬리표만잘라냄
base = re.sub(
r'\.(?:[a-z]{2,3}(?:-[A-Za-z0-9]+)*)(?:-orig)?$',
''
,
    base_no_ext,
    flags=re.I
)
```

전체 함수 모양 참고(나머지는 동일): 

python 

코드 복사 

```
defextract_title_key(path):
    name = os.path.basename(path)
    base_no_ext = os.path.splitext(name)[0]
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

33/65 

25. 11. 30. 오후 1:39 

 

후 1:39 1:39 프로그램 - 유튜브 배치파일 수정 _`#`_ ⬇ `여기만 위 패턴으로 교체 base = re.sub( r'\.(?:[a-z]{2,3}(?:-[A-Za-z0-9]+)*)(?:-orig)?$', '' , base_no_ext, flags=re.I ) m = re.match(r'^(?P<title>.+?)\s*\[[^\]]+\]$', base) raw_title = (m.group('title') if m else base).strip() return normalize_title(raw_title)` 

이렇게 바꾸면 예시처럼 

`[...].ko.vtt` / `[...].ko-orig.vtt` → **동일 제목 키** 로 합쳐지고, 그중 **더 긴 텍스트** (보통 수동/정식 자막)가 대표로 채택돼. 

나머지 파일( `Scirpt.bat` )은 그대로 쓰면 돼. 

테스트 중 또 다른 꼬리표가 보이면(예: `.ko-auto` , `.ja-orig` ) 위 정규식이 이미 포괄적으 로 처리하니 추가 수정 없이 잘 묶일 거야. 

**==> picture [489 x 39] intentionally omitted <==**

## 왜 이렇게 두개의 영상을 업로드했지? 

50s 동안 생각함 

## 두 개가 **영상이 아니라 자막 트랙** 이야. 

`yt-dlp` 옵션을 `--write-subs` (수동자막) + `--write-auto-subs` (자동자막) 둘 다 켜둬서, 같은 제목에 대해 

- `…ko.vtt` → 한국어 자막(대개 자동/번역 포함) 

- `…ko-orig.vtt` → 원본 트랙을 구분하려고 붙인 “orig” 버전 

## 이렇게 **두 파일이 내려온 것** 이야. 

내가 준 `Youtube.py` 는 이런 경우 **같은 제목으로 묶어서 하나로 통합** 하고, 더 **긴 쪽** 을 대 표로 선택한 뒤 성공 시 원본 자막들을 지워버리므로, 결과물은 1개로 정리돼. 

만약 “애초에 하나만 받자”(수동 우선, 없으면 자동)로 바꾸고 싶으면 `Scirpt.bat` 의 다운 로드 구간을 아래처럼 바꿔줘. (이러면 `ko` / `ko-orig` 중복 자체가 사라짐) 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

34/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

코드 복사 

bat 

```
for /f "usebackq delims=" %%U in ("urls.txt") do (
```

```
  echo [DL] %%U
```

```
  rem 1) 수동자막(ko) 먼저시도
```

```
  yt-dlp %COOK% --write-subs --sub-langs "ko" --skip-download --convert-subs vtt
```

```
  rem 2) 방금파일이하나도없으면자동자막(ko)로대체
```

```
  dir /b "*.vtt" >nul 2>nul || yt-dlp %COOK% --write-auto-subs --sub-langs "ko" -
```

- `)` 

 

 

## 요약 

- 두 파일이 생긴 이유: **자막을 수동+자동 둘 다 받도록 설정** 되어 있어서. 

- 현재 파이썬 스크립트는 같은 제목이면 **자동 통합** 처리. 

- 중복 자체를 막고 싶으면 위 **배치 패치** 를 적용하면 돼. 

**==> picture [206 x 193] intentionally omitted <==**

영상에서도 이렇게 동일한 두개의 파일을 업로드했어 무 슨이유지? 

## 19s 동안 생각함 

형, 저건 “중복 업로드”라기보다 **같은 주제의 두 버전** 을 올린 케이스일 가능성이 커. 스샷 보면 길이가 서로 달라(예: 32:35 vs 22:45). 보통 이유는 이 중 몇 가지야: 

**1. 풀버전 + 요약본** : 시청 유지율(AVD) 실험용으로 긴 버전과 짧은 버전을 같은 제목으 로 올림. 

**2. 라이브 재업로드** : 라이브 전체본(길다) + 사후 편집본(불필요 구간 컷). 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

35/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

**3. 오류 수정/갱신본** : 실수·음악저작권·오타 수정해서 같은 제목으로 다시 올리고, 기존 영상은 메트릭 때문에 지우지 않음. 

**4. 알고리즘 테스트** : 제목은 고정하고 길이·인트로·콜투액션만 다르게 해서 CTR/AVD 비 교. 

**5. 자막/해상도 차이** : 자막·챕터·자막언어 구성 다른 버전(덜 흔하지만 있음). 

**6. 리메이크** : 같은 주제를 최신 라이브러리/코드로 다시 찍음(길이 변동 큼). 

## **빠른 확인법** 

- 두 영상의 **업로드 날짜/설명/고정댓글** 비교(“수정본”, “요약본” 표기 여부). **챕터(타임스탬프)** 유무/구성 비교. 

- **영상 ID** 가 다르면 재업로드/별도 버전. 

- 코딩 영상이면 **데모/코드 링크** 가 바뀌었는지 확인. 

## **파이프라인 쪽 처리는 이미 OK** 

우리가 만든 `Youtube.py` 는 **같은 제목이면 하나로 통합** 하고, **더 긴 텍스트(대개 더 온전 한 자막)**를 대표로 채택해. 그래서 이런 이중 업로드가 있어도 결과물은 1개로 정리돼. 

원하면 이런 **로그** 를 남겨서 왜 어떤 버전을 대표로 골랐는지 한눈에 보이게도 해줄 수 있 어: 

- 같은 제목으로 묶인 영상들의 `제목키` , `영상ID` , `길이(자막 길이)` , `선택된 대표` 를 `Youtube_log.txt` 에 표 형식으로 기록. 

필요하면 그 로그 기능 패치도 바로 붙여줄게. 

**==> picture [489 x 39] intentionally omitted <==**

왜냐하면 많은 영상을 다이렉트로 스크립트 전환에 적용 시키다보니 

유튜브에서 에러형태로 인식을 하는 경우가 있어서 그래 

48s 동안 생각함 

형 말이 맞아. 다이렉트로 많이 때리면 유튜브가 비정상 트래픽으로 오인할 수 있어. 그래서 **예의 있게(throttling + 랜덤 딜레이 + 재시도·백오프)** 돌도록 배치만 살짝 강화해 줬어. 파이썬 파일은 그대로(제목 통합·중복 제거·4만자 분할 유지). 

파일명 고정: 배치 = **Scirpt.bat** , 파이썬 = **Youtube.py** . 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

36/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

## **Scirpt.bat (업데이트: 안전 다운로드 모드 탑재)** 

bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브목록처리(폴더지정) - TXT 통합/분할(세이프모드)
rem === 기본경로 ===
set "ROOT=%~dp0"
echo 저장할폴더이름을입력하세요. (빈칸이면날짜_시각으로자동생성)
set "FOLDER="
set /p FOLDER=폴더명:
if "%FOLDER%"=="" (
  for /f %%i in ('powershell -NoProfile -Command "(Get-Date).ToString(\"yyyy-MM-dd
)
set "DEST=%ROOT%%FOLDER%"
if not exist "%DEST%" mkdir "%DEST%"
echo.
echo [안내] URL을한줄씩붙여넣고 Enter 하세요.
echo        아무것도입력하지않고 Enter 하면다운로드를시작합니다.
echo.
pushd "%DEST%"
```

```
if exist urls.txt del /f /q urls.txt
```

```
:INPUT_LOOP
set "URL="
set /p URL=URL:
if "!URL!"=="" goto START
echo !URL!>>urls.txt
goto INPUT_LOOP
:START
if not exist urls.txt (
  echo [오류] URL이없어작업을종료합니다.
  popd
  pause
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

37/65 

25. 11. 30. 오후 1:39 

 

프로그램 - 유튜브 배치파일 수정 

```