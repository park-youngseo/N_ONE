25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

# ===================================== ======= # 엔오원 니즈 수집 크롤러 v2.4 # 그래비티 생성용 완성본(복붙용) # ===================================== ======= 

# 생성 규칙 # 1) 아래에 정의된 파일명을 그대로 생성한다. # 2) cleaner.py / tagger.py / patterns.py / post.py / saver.py / logger.py 는 수정 금지. # 3) fetcher.py / main.py 만 그래비티가 자동 생성. # 4) output / logs 폴더 자동 생성. # 5) 엔진 코어는 절대 변형하지 말고 그대로 만든다. 

# 

===================================== ======= 

# 1. cleaner.py # ===================================== ======= 

[cleaner.py 파일 내용 시작] 

import re 

def clean_text(text: str) -> str: if not text: return "" # URL 제거 text = re.sub(r"http[s]?://\S+", "", text) # HTML 태그 제거 text = re.sub(r"<[^>]+>", "", text) # 광고/홍보성 패턴 제거 text = re.sub(r"(문의|카톡|오픈채팅|insta|인스타|DM주 세요|구독|좋아요|체험단)", "", text) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

1/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 # 이모지 제거  text = re.sub("[\U00010000-\U0010ffff]", "", text) # 반복 문자 2개 제한 text = re.sub(r"(.)\1{2,}", r"\1\1", text) # 양쪽 공백 정리 text = text.strip() return text 

[cleaner.py 파일 내용 끝] 

===================================== ======= 

# 2. tagger.py # 

===================================== ======= 

[tagger.py 파일 내용 시작] 

from patterns import * 

def apply_tags(text: str) -> str: tags = [] 

for word in AUDITION_WORDS: if word in text: tags.append("<심리_오디션>") 

for word in CASTING_WORDS: if word in text: tags.append("<심리_드라마캐스팅>") 

for word in START_WORDS: if word in text: tags.append("<심리_시작>") 

if tags: 

return text + " " + " ".join(set(tags)) return text 

[tagger.py 파일 내용 끝] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

2/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

===================================== ======= # 3. patterns.py # 

===================================== ======= 

[patterns.py 파일 내용 시작] 

# 스탭 모집 제외 STAFF_HINTS = [ "스탭", "스태프", "촬영 도와주실", "현장 도와주실", "스 탭모집" ] 

# 광고 제거 AD_HINTS = [ "문의", "카톡", "카카오", "체험단", "DM주세요", "홍보" ] 

# 심리 태그용 AUDITION_WORDS = ["오디션", "프로필", "연기영상", "캐 스팅디렉터"] CASTING_WORDS  = ["드라마", "배역", "출연"] START_WORDS    = ["시작", "입문", "고민"] 

# 제거 규칙 

def is_staff_post(text: str) -> bool: 

return any(word in text for word in STAFF_HINTS) 

def is_ad_post(text: str) -> bool: 

return any(word in text for word in AD_HINTS) 

[patterns.py 파일 내용 끝] 

## 

===================================== ======= # 4. post.py 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

3/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

===================================== 

======= 

[post.py 파일 내용 시작] 

from cleaner import clean_text from tagger import apply_tags from patterns import is_staff_post, is_ad_post 

def process_post(post: dict) -> dict: 

title = clean_text(post.get("title", "")) body = clean_text(post.get("body", "")) comments = clean_text(post.get("comments", "")) 

merged = f"{title} {body} {comments}".strip() 

- if not merged: 

- return None 

- if is_staff_post(merged): 

- return None 

- if is_ad_post(merged): 

- return None 

merged = apply_tags(merged) 

return { 

- "title": title, 

- "body": body, 

- "comments": comments, 

- "fulltext": merged 

[post.py 파일 내용 끝] 

## 

===================================== 

======= 

# 5. saver.py 

===================================== 

======= 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

4/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

[saver.py 파일 내용 시작] 

import csv import os 

def save_csv(filename: str, rows: list): 

os.makedirs("output", exist_ok=True) 

filepath = os.path.join("output", filename) 

with open(filepath, "w", encoding="utf-8-sig", newline="") as f: 

writer = csv.writer(f) 

writer.writerow(["title", "body", "comments", "fulltext"]) 

for r in rows: 

writer.writerow([r["title"], r["body"], r["comments"], r["fulltext"]]) 

return filepath 

[saver.py 파일 내용 끝] 

===================================== ======= 

# 6. logger.py 

===================================== ======= 

[logger.py 파일 내용 시작] 

import os, datetime 

def log(message: str): 

os.makedirs("logs", exist_ok=True) 

now = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S") 

with open("logs/crawler.log", "a", encoding="utf-8sig") as f: 

f.write(f"[{now}] {message}\n") 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

5/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

[logger.py 파일 내용 끝] 

===================================== ======= 

# 7. fetcher.py # 

===================================== ======= 

[fetcher.py 파일 내용 시작] 

import requests from bs4 import BeautifulSoup import time import random from post import process_post from logger import log 

headers = { "User-Agent": "Mozilla/5.0 (Windows NT 10.0)" } 

def fetch_generic_board(url: str) -> list: posts = [] log(f"크롤링 시작: {url}") 

page = 1 while True: page_url = f"{url}?page={page}" 

r = requests.get(page_url, headers=headers, timeout=10) 

if r.status_code != 200: 

break 

soup = BeautifulSoup(r.text, "html.parser") items = soup.select("a")  # 모든 링크 우선 수집 

if not items: break 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

6/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

for a in items: title = a.text.strip() href = a.get("href", "") if not href: continue if href.startswith("/"): href = url.split("/")[0] + "//" + url.split("/") [2] + href # 상세 페이지 try: d = requests.get(href, headers=headers, timeout=10) ds = BeautifulSoup(d.text, "html.parser") body = ds.get_text(" ", strip=True) comments = "" row = { "title": title, "body": body, "comments": comments } processed = process_post(row) if processed: posts.append(processed) except: continue page += 1 time.sleep(random.uniform(1, 3)) except Exception as e: log(f"오류 발생: {e}") log(f"수집 완료: {len(posts)}개") return posts 

[fetcher.py 파일 내용 끝] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

7/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

===================================== ======= # 8. main.py # ===================================== ======= 

[main.py 파일 내용 시작] 

from fetcher import fetch_generic_board from saver import save_csv from logger import log 

def ask_user(): 

print("=============================") print("  엔오원 니즈 수집 크롤러 v2") 

print("=============================\n") 

site_name = input("[1] 사이트 이름 입력: ").strip() if not site_name: print("사이트 이름은 필수입니다.") exit() 

print("\n[2] 크롤링할 게시판 URL 입력") print("입력이 끝났으면 빈 줄에서 Enter") urls = [] while True: u = input("> ").strip() if u == "": break urls.append(u) 

if not urls: print("URL이 없습니다.") exit() 

return site_name, urls 

def run(): 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

8/114 

25. 11. 30. 오후 1:06 
