# How I Made a Score System in Unreal Engine 5.6 in 7 Minutes
## [메타데이터]

| 항목 | 내용 |
|------|------|
| **채널** | Unreal Fusion |
| **영상 ID** | tjA5mO-JZFg |
| **영상 URL** | https://www.youtube.com/watch?v=tjA5mO-JZFg |
| **카테고리** | 10_Unreal_Engine |
| **파트** | 1 / 1 |
| **변환일** | 2026-05-04 |

---

## [원본 자막 전문 — 무손실 (Part 1)]

> **[AX 하네스 준수]** 요약 및 압축 절대 금지. 원본 자막 텍스트 전문 보존.

Hi everyone. In this tutorial, I want to
show you how to make score system like
this.
The score is zero. When I touch the
coin,
the score is increased by one.
Okay.
Okay. First thing first, open the
engine.
Go to game. Insert person. Rename the
project score system. Once it's done,
hit create.
Now let's go here and create
user interface widget for the score. Ren
named it score UI
UI. Go here inside the score UI.
Make it right there.
Go to the palette and choose overlay.
Now inside the overlay let's search for
is on horizontal boxes
and this horizontal box we have to make
two text
first one.
Okay fine. Press text
renamed
score.
Now select the horizontal box and go for
search for translation.
Now select the second one.
Go to content and create a binding.
After that we have to go to the event
graph this corner and cast to curtain
character
in this place. Get play
character.
drag it here as third person character.
Let's promote to a variable.
Let's return here
and drag this one. Get
now hit compile. Save all. Return right
here to third person blueprint third
person character.
In this place we have to create a
variable score.
The type of this variable is integer.
After that we have to display my widget
in the viewport.
This one. To do that in the same
blueprint class we have to search for
begin event begin play
widget
my widget class score UI after that add
to view port
I think this step is fine
now let's return here
to my score UI
and get the score value. Get the score.
After getting the score, get here the
return value.
Okay, compile and save. Now let's go
here
from the content drawer. Add go to fab.
Search for a coin.
Uh, choose anyone.
Let's choose this coin for this
tutorial.
Add to project.
Okay, my coin is fine. Now let's go here
and create new blue blueprint class to
type actor
and this class coin
score
when the player touches this coin the
score is
increased by one. And now from this one,
let's search for box collision.
And this box collision,
we have to go here and add my coin. To
do that, we have to go here to the fab
coin
static mesh and drag this coin inside
the vertical box inside the box
collision. Now select this box
collision. Let's go down in this corner
on component begin overlap.
Now
let's cast it to the first character to
the
person character
for the object other object and this
from the first person character here.
Let's search for get score
here. Now let's add
+ one
one to this code. Now let's set the
score.
Build the
score.
Find make it right there. Okay. After
doing that we have to destroy the actor.
destroy my coin after touching it.
Now save all. I think it's done. After
doing that, we have to go here
and drag the coin
here.
Press alt
and duplicate this coin.
And let's see it works or not. Uh go
here. This is my third score. The first
score is zero.
Okay.
One,
two.
I want to get this coin, but the problem
is
I Okay, but no, it's not happen. Okay,
thank you guys for watching. This is all
for this tutorial. Uh, see you later on
the next tutorial. Subscribe. Subscribe
for the channel for more tutorials like
this.

---

*변환 완료: 2026-05-04 | 원본 무손실 전문 | AX 표준 적용*
*다음 파트: 없음 (최종)*


---
## 🔗 관련 시너지 지식
- **핵심 하네스**: [Global_Scenario_Harness_21.md](../../Global_Scenario_Harness_21.md)
- **클러스터**: Creative Production
