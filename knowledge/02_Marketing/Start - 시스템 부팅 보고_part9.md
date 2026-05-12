
 

- README.txt 

- .env.example requirements.txt seed_keywords.csv mvp_collect.py 

## **어떻게 쓰면 되냐면** 

**1.** 파이썬 3.10+ 설치 후, 패키지 설치 

bash 코드 복사 `pip install -r requirements.txt` API 키 넣기 bash 코드 복사 

**2.** API 키 넣기 

```
cp .env.example .env
# .env 열어서키입력
# YOUTUBE_API_KEY=...
# NAVER_CLIENT_ID=...
# NAVER_CLIENT_SECRET=...
```

**3.** 수집 실행 → 결과 CSV 생성 

bash 코드 복사 

```
python mvp_collect.py --seed seed_keywords.csv --months 12 --out results.csv
```

- 출력: `results.csv` (컬럼: `keyword, source, yyyymm, count` ) 

- 예: `오디션, youtube, 2025-08, 42` 

## **포함된 시드(Non-입시, 20개 예시)** 

오디션, 연기학원, 카메라연기, 프로필사진, 독백연습, 오디션영상, 카메라 테스트, 연극/드 라마/뮤지컬 오디션, 표정연기, 감정이입, 대사 암기, 즉흥연기, 리딩, 모노로그, 워크샵, 씬 스터디, 단편영화, 웹드라마 

원하면 바로 **FastAPI 대시보드 / Postgres 저장 / 감정·의도 라벨 집계** 까지 붙여서 확장형 도 만들어줄 수 있어. 우선 이 MVP로 실제로 **검색→집계→결과 뽑기** 부터 확인해보고, 잘 돌아가는 게 보이면 다음 단계로 가자. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

43/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

나는 프로그램 제작 지식대해 아무것도 모르는 상태야 초 등학생도 이해할수 있게 차례차례 하나씩 만들어볼수 있 게 해줘 

좋아 형! 완전 처음부터, **초등학생도 할 수 있게** 진짜 쉬운 단계로 하나씩 가자. (지금 당장 쓸 수 있게 파일도 묶어놨어요 👉 ActorTrend_MVP_v0_1.zip) 

## **0단계: 이게 뭔지 한 줄 설명** 

이건 “배우 관련 단어”를 유튜브·네이버 블로그에서 **한 달에 몇 번 언급됐는지** 숫자로 모아 표(CSV)로 보여주는 **작은 프로그램** 이야. 

## **1단계: 파일 받기** 

**1.** 위에 있는 **ActorTrend_MVP_v0_1.zip** 을 눌러서 저장해. 

**2.** 압축을 풀면 이런 파일이 있어: 

   - `README.txt` (설명서) 

   - `.env.example` (열쇠 넣는 종이) 

   - `requirements.txt` (필요한 도구 목록) 

   - `seed_keywords.csv` (검색할 단어 목록) 

   - `mvp_collect.py` (프로그램 본체) 

## **2단계: 파이썬 설치 (컴퓨터에 없으면)** 

**윈도우** : 구글에서 “파이썬 다운로드” 검색 → Python 3.10 이상 설치 → 설치 중 “Add Python to PATH” 체크. 

- **맥** : App Store 말고 python.org에서 Python 3.10 이상 설치. 

## **3단계: 프로그램이 쓸 도구 설치** 

**1.** 압축 푼 폴더로 들어가서, 

**2.** 아래 중에 내 컴퓨터에 맞는 한 줄만 복사해서 실행해. 

   - **윈도우** 

nginx 코드 복사 `pip install -r requirements.txt` 

## **맥** 

nginx 

코드 복사 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

44/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
pip3 install -r requirements.txt
```

## **4단계: 열쇠(키) 넣기** 

이 프로그램은 **유튜브/네이버** 에서 데이터를 가져오려면 각각의 **열쇠(API 키)**가 필요 해. 

**1.** 폴더 안에 있는 **.env.example** 파일을 **복사해서 이름을** `.env` **로 바꾸기** 

**2.** 메모장으로 `.env` 를 열고, 이렇게 채우기: 

ini 코드 복사 

```
YOUTUBE_API_KEY=여기에_유튜브_API키
NAVER_CLIENT_ID=여기에_네이버_클라이언트ID
NAVER_CLIENT_SECRET=여기에_네이버_시크릿
```

유튜브 API 키와 네이버 키가 당장은 없으면, **빈칸으로 둬도 실행은 돼요** (그럼 해당 쪽 데이터는 0으로 나올 뿐). 

## **5단계: 바로 실행하기** 

아래 중 내 컴퓨터에 맞는 한 줄만 실행해. 

## **윈도우** 

css 코드 복사 

```
python mvp_collect.py--seed seed_keywords.csv--months12--out results.csv
```

**==> picture [457 x 64] intentionally omitted <==**

**----- Start of picture text -----**<br>
 <br>맥<br>css 코드 복사<br>**----- End of picture text -----**<br>


```
python3 mvp_collect.py--seed seed_keywords.csv--months12--out results.csv
```

  

실행이 끝나면 같은 폴더에 **results.csv** 가 생겨요. 엑셀로 열면 이런 표가 보여요: 

|**keyword**|**source**|**yyyymm**|**count**|
|---|---|---|---|
|오디션|youtube|2025-08|42|
|오디션|naver_blog|2025-08|17|



https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

45/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

**keyword source** 

**count** 

**yyyymm** 

… … … … 

## **6단계: 이 숫자를 어떻게 보냐?** 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

46/46 
