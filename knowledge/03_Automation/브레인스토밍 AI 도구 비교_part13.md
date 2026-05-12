
운영비 최소, 도입 1일 내 사용 가능 

Google Sheets + Google Calendar 연동(선택) 

명령어 기반: `/add` , `/move` , `/cancel` , `/planweek` , `/report` 

## 🧩 **시스템 구성(요약)** 

**1. 대화 엔진** : ChatGPT (이 문서의 프롬프트·명령어 사용) 

**2. 데이터 저장소** : Google Sheets(스케줄 DB) 

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

66/72 

25. 11. 30. 오후 1:02 

브레인스토밍 AI 도구 비교 

 

**3. 캘린더** : Google Calendar(선택) — Apps Script로 양방향 동기화 

**4. 입력 채널** : ChatGPT 대화창(형이 명령), 추후 카카오 채널 연동 가능 

## 📚 **데이터 스키마(스프레드시트)** 

스프레드시트 이름: `ENONE_Schedule` 

## 시트 탭: **Schedule** 

**==> picture [446 x 521] intentionally omitted <==**

**----- Start of picture text -----**<br>
|||||
|---|---|---|---|
|Column|Type|예시|설명|
|A: id|text|EVT-2025-10-14-001|고유ID(자동|
|생성)|
|B: title|text|메커니즘 Class (TEAM A)|일정 제목|
|C: date|date(YYYY-|2025-10-15|날짜|
|MM-DD)|(Asia/Seoul)|
|D: start_time|text(HH:MM)|15:00|시작 시각|
|E: end_time|text(HH:MM)|17:00|종료 시각|
|F: location|text|강의실 1|장소/링크|
|G: category|enum|수업/촬영/편집/홍보/행정|업무 종류|
|H: priority|enum|A/B/C|우선순위(A|
|가 최상)|
|I: notes|text|체험 2석|메모|
|J: status|enum|pending/confirmed/moved/canceled|상태|
|K:|text|...|캘린더 이벤|
|gcal_event_id|트 ID(자동|
|기입)|
|L:|datetime|...|마지막 수정|
|last_updated|(자동)|

**----- End of picture text -----**<br>


**규칙** : 챗봇은 기본적으로 `pending` 으로 기록 후 확인 시 `confirmed` 로 전환. 

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

67/72 

25. 11. 30. 오후 1:02 

브레인스토밍 AI 도구 비교 

 

## 💬 **챗봇 명령어(형이 그대로 사용)** 

## **1) 일정 추가 —** `/add` 

```
/add [제목]; [날짜]; [시작~종료]; [카테고리]; [우선순위]; [장소]; [메모
```

  

## **예시** 

```
/add 메커니즘 Class (TEAM A); 2025-10-15; 15:00~17:00; 수업; A; 강의실1
```

  

## **2) 일정 이동 —** `/move` 

```
/move [id 또는 제목 일부]; [새 날짜]; [새 시작~종료]
```

예) 

```
/move EVT-2025-10-14-001; 2025-10-16; 17:00~19:00
```

## **3) 일정 취소 —** `/cancel` 

```
/cancel [id 또는 제목 일부]; [사유]
```

## **4) 주간 계획 생성 —** `/planweek` 

```
/planweek
```

```
목표: (예) 릴스1+카드뉴스1, 체험2회, 등록1~2명
```

```
고정블록: 수업(수/토 2h), 촬영(화 2h), 카드뉴스(목 2h), 행정(월 1h)
제한: 수업>콘텐츠>행정 우선. 저녁 7시 이후 수업만 가능.
```

→ 챗봇이 빈 시간대 자동 배치 + 제안(옵션 2안) + `/add` 명령어로 변환까지 출력. 

## **5) 주간 리포트 —** `/report` 

```
/report [YYYY-MM-DD 주 시작일]
```

## **출력 예시** 

총 일정 ○건(수업 ○, 촬영 ○, 홍보 ○, 행정 ○) 

완료율 ○%, 이동/취소 ○건 

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

68/72 

25. 11. 30. 오후 1:02 

브레인스토밍 AI 도구 비교 

 

제안: “촬영을 화요일 고정하면 겹침 2건 감소 예상” 

## 🧠 **챗봇용 시스템 프롬프트(복붙)** 

**역할** : 당신은 엔오원(none)의 스케줄 전용 비서다. 한국어로 답한다. 모든 시간 대는 Asia/Seoul. 명령어 기반으로 일정을 생성/이동/취소하고, 구조적 응답을 제공한다. 

**규칙** : 

**1.** “출력은 항상 다음 순서: 요약 → 상세(표 또는 코드블록) → 다음 행동 제 안.” 

**2.** “ `/add` 요청은 누락된 필드가 있으면 질문 1회 후 기본값으로 채워서 제안 안 A/B 출력.” 

**3.** “충돌 탐지: 같은 날짜·시간 겹치면 경고 + 자동 재배치 제안(2안).” 

**4.** “포맷: `YYYY-MM-DD` 날짜, `HH:MM` 24시간.” 

**5.** “카테고리 {수업, 촬영, 편집, 홍보, 행정} 중 선택. 우선순위 {A,B,C}.” 

**6.** “시트 저장은 ‘명령어 변환 텍스트’도 함께 제공해 사용자가 기록하기 쉽게 한다.” 

**7.** “리포트에는 합계/이동·취소 건수/완료율/개선 제안을 포함.” 

## 🧰 **Google Apps Script — Sheets↔Calendar 동기화 (선택)** 

**1.** 스프레드시트 `확장 프로그램 → 앱스 스크립트` 열기, 아래 코드 붙여넣기 → 저 장 → 권한 승인. 

```
constSHEET = 'Schedule';
```

 

```
constCALENDAR_ID = '여기에_형의_캘린더_ID'; // 캘린더 설정에서 ID 복
```

```
functiononEdit(e) {
constsh = e.range.getSheet();
if (sh.getName() !== SHEET) return;
constr = e.range.getRow();
constrow = sh.getRange(r, 1, 1, 12).getValues()[0];
const [id, title, date, start, end, location, category, priority, no
```

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

69/72 

25. 11. 30. 오후 1:02 브레인스토밍 AI 도구 비교 `if (!title || !date || !start || !end) return;`   `if (status === 'canceled' && evId) { deleteEvent(evId); sh.getRange if (evId) updateEvent(evId, row); else createEvent(r, row); sh.getRange(r,12).setValue(new Date()); } function createEvent(rowIndex, row){ const cal = CalendarApp.getCalendarById(CALENDAR_ID); const [id, title, date, start, end, location, category, priority, no const startDt = toDateTime(date, start); const endDt = toDateTime(date, end); const event = cal.createEvent(title, startDt, endDt, {location: loca const sh = SpreadsheetApp.getActive().getSheetByName(SHEET); const eid = event.getId(); sh.getRange(rowIndex, 11).setValue(eid); } function updateEvent(evId, row){ const cal = CalendarApp.getCalendarById(CALENDAR_ID); const ev = cal.getEventById(evId); if (!ev) return; const [id, title, date, start, end, location, category, priority, no ev.setTitle(title); ev.setTime(toDateTime(date, start), toDateTime(date, end)); if (location) ev.setLocation(location); ev.setDescription(`[${category||''}|${priority||''}] ${notes||''}`. } function deleteEvent(evId){ const cal = CalendarApp.getCalendarById(CALENDAR_ID); const ev = cal.getEventById(evId); if (ev) ev.deleteEvent(); } function toDateTime(date, time){ const [h,m] = time.split(':').map(Number); const d = new Date(date); d.setHours(h, m, 0, 0); return d; } // 수동 동기화 (초기 일괄 업로드) function syncAll(){ const sh = SpreadsheetApp.getActive().getSheetByName(SHEET); const rows = sh.getDataRange().getValues(); for (let r=2; r<=rows.length; r++){` 

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

70/72 

25. 11. 30. 오후 1:02 

 

 

브레인스토밍 AI 도구 비교 

```
constrow = sh.getRange(r,1,1,12).getValues()[0];
const [id, title, date, start, end, location, category, priority,
if (!title || !date || !start || !end) continue;
```

```
if (status==='canceled' && evId){ deleteEvent(evId); sh.getRange(
if (evId) updateEvent(evId, row); elsecreateEvent(r, row);
sh.getRange(r,12).setValue(newDate());
}
```

## ✅ **온보딩 체크리스트 (10분)** 

**1.** 구글 스프레드시트에 `ENONE_Schedule` 파일 만들고 상단 스키마대로 **Schedule** 탭 생성 

**2.** `확장 프로그램→앱스 스크립트` 열고 상단 코드 전체 붙여넣기 → `CALENDAR_ID` 입력 → 저장/권한 승인 

**3.** 테스트 행 1개 추가: 

`EVT-TEST-001 | 메커니즘 Class (TEAM A) | 2025-10-15 | 15:00 | 17:00 | 강`   → 셀 수정하면 캘린더에 자동 생성되는지 확인 4) ChatGPT에 이 문서의 **시스템 프롬프트** 붙여넣고, 아래 명령어 실행 

```
/planweek
```

```
목표: 릴스1+카드뉴스1, 체험2
고정블록: 수업(수/토 2h), 촬영(화 2h), 카드뉴스(목 2h), 행정(월 1h)
제한: 수업>콘텐츠>행정. 19시 이후 수업만 가능.
```

→ 출력된 제안안에서 마음에 드는 슬롯들을 `/add ...` 로 확정 

## 🧪 **테스트 시나리오** 

**1. 추가** : `/add 촬영; 2025-10-16; 13:00~15:00; 촬영; B; 스튜디오; 유튜브 숏츠` 

**2. 겹침 탐지** : 같은 시간대에 수업 추가 → 재배치 제안 2안 확인 

**3. 이동** : `/move 촬영; 2025-10-17; 14:00~16:00` 

**4. 취소** : `/cancel 메커니즘 Class; 사유: 수강생 요청` 

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

71/72 

25. 11. 30. 오후 1:02 

브레인스토밍 AI 도구 비교 

 

**5. 리포트** : `/report 2025-10-13` 

## 📎 **수동(연동 없이) 사용 팁** 

연동이 번거로우면 초기에 **시트 없이** 도 사용 가능: 챗봇이 표 형태로만 출력 → 형이 복붙 기록 

나중에 여유 생기면 시트/캘린더 동기화만 추가하면 됨 

## 🔒 **가드레일(운영 규칙)** 

모든 시간·날짜는 **Asia/Seoul** 

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

72/72 
