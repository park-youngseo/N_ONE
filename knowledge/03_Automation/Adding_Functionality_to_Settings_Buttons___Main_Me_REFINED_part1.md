# Adding Functionality to Settings Buttons | Main Menu in Unreal Engine 5.5.4 (Part 4)
## [메타데이터]

| 항목 | 내용 |
|------|------|
| **채널** | Unreal Fusion |
| **영상 ID** | MVSaFij5wJg |
| **영상 URL** | https://www.youtube.com/watch?v=MVSaFij5wJg |
| **카테고리** | 10_Unreal_Engine |
| **파트** | 1 / 1 |
| **변환일** | 2026-05-04 |

---

## [원본 자막 전문 — 무손실 (Part 1)]

> **[AX 하네스 준수]** 요약 및 압축 절대 금지. 원본 자막 텍스트 전문 보존.

Hello everyone. In this part, we want to
create the button
functionalities. We have we have to work
on programming all the buttons of the
shadows, texture, frame rates,
persp. Okay. Uh first of all, we have to
go to the graph right here.
And we have to
create add functionality to those
buttons.
Okay, let's begin with the shutters. Go
to this di select
low. Go
down. First one shadows
low. Okay. And go right here.
Shadow
medium. Go to the event on click
it shadow
high shadow
epic and
finally shadow
cinematic. Okay, let's add shadow
quality. First of all, we have to get
the game user settings.
From the return
value we have to set shadow
quality. Okay. For the shutter quality
for the quality is begins by zero to
four. Zero for the low axis and low
quality. One for the
medium sorry
make
two for the highest and three for the
apric and four for the seimatic. Okay
let's begin. Okay it's set by default to
value to zero.
Ctrl C. Ctrl
V. Okay. Ctrl V. And
finally, Ctrl
VR
V. Okay. Drag all of
this. And now let's apply the setting
for all of this.
apply
settings. Okay. For the
target, let's make it like
this. This value is
one. Ctrl C. Ctrl V.
Okay. Make it like
this. Ctrl C. Ctrl
V. You can make it like this.
This is
two,
three, and finally
four. Control C. Control V.
Okay, after doing this, make sure to
then check this
one.
Okay. Okay. Now, let's command all of
those. Select all and press
C. And set this one to shutter
quality. Okay. And let's see it works or
not. Hit compile and save
all.
Now let's add with twister
functionality. Okay.
We have to switch between uh those one
settings and the main menu. We go here
to my widget
switcher active widget by index set to
zero. We have to switch between this and
the setting
menu. Okay. To do that, we have to go
here to the
graph. Select the back button.
Where is the back? Okay, let's select
from
this. Let's go here. Select second
overlay. And let's go
down. This is back button. And the other
one is sitting button.
Let's go here. Select this
thing. We have to switch between
settings and main menu. I'll click
it.
Okay. Now, let's set uh set my widget
widget. Select this one and drag it from
here to here to this place.
and set active
widget. Active widget by
index. The first one when I select the
back button, it's switch to the main
menu. Set index to
zero. Okay. Ctrl C and Ctrl
D. The index of this thing is one. Let's
set it to
one. Okay. Compile and save. And let's
test
it.
Then settings. When I press the back
button, it returns the main menu. It
works
fine. Okay. Let's return
here. Let's comment this which is
switcher by index.
by
index. Okay. And now let's test my
shadow quality. It's for
some. Let's go
here. Let's go to
snadow quality to low. Return right
here.
Go to start
game. When we are we are see we have no
shadows. It works
fine. Okay. When you go here the
scalability it's
slow.
Now let's
move to texture
quality. Go right here. Where is my
texture?
Maybe this horizontal
box. This
one text
low. Go
down. Okay.
texture
medium texture
high texture opic and texture
cinematic. Let's do the same process but
in this case we have to set uh set the
texture quality.
game
settings. I get
game game user
settings. From this one, we have to set
the
text quality.
Okay. From this we have to apply
settings.
Okay. Make it like
this. Now we have to do is to duplicate
this one to the
others. Okay. and change the index
Okay. Hello
For mediums this one and take a value
one,
two,
three, and finally
four. Okay, comment all of this and make
it texture quality.
Compile
antiode make this one right here. What
is the
[Music]
problem? Now let's go to design. Let's
move to frame rates per
second. Go
here and select my frame rates per
second. Okay, let's begin by 30
fps, 60 fps, 90
fps, and finally unlimited FPS.
This one unedited
but okay. Get
game user settings. Now we have to make
the frame
[Music]
rates.
Set frame rate is
limit. Okay. The first one is uh 30 fps.
We have to set this one 30 fps.
Make it the target like
this. And finally,
apply
settings. Make sure this one is
unchecked.
Ctrl Ctrl
V.
Okay. Drag all of
this. Now this one is 60 fps.
fps and finally set this to zero to make
it
unlimited. Select all of this. Command
it. This is frame
rate second.
Select all of this and make it right
there.
Okay. Uh make this uncheck it. I forget
to add
it. Okay.
Fine. Compile and see all. Let's go to
designer.
You have to make this one the final one.
The
molds. Select the first
button. I'll click
it. And the second button is the full
screen button. Go
down. I'll take
it. Let's begin for the first one.
and get
user game
settings. Okay. From
this set full screen
mode for the first one we have to add to
window
mode and finally
apply settings.
I check this
[Music]
one. Ctrl C and Ctrl
D.
Okay, this is one to full
screen. Come on out of this and set this
one modes.
Save all and return here. The final one
is apply best settings. For this button,
we have to use your benchmark your
computer and set when you click it, it
set the best uh quality, best settings
for your
computer. Okay.
To do that we have to get user game
settings game user
settings from this one. Let's run a
benchmark your hardware
benchmark and
apply
hardware benchmark results.
Connect right
here. Okay. Ctrl C and command of this.
Apply best game.
settings. Okay, fine. Let's test
out apply.
Compile and save all. Let's go here. Now
to settings. I want to set for example
set quality to high, texture low, and
frame rates for,000 second. And let's
set to full
screen. Back. Start
game. Okay, it's fine. Let's just let's
go to this
liability and look at this shadows high
texture.
We have to make the
dough. Okay. Thank you guys for watching
this all for this
tutorial. On the next
tutorial, when I apply enabling buttons
and disabling button, for example, when
I click below, it's disabled
automatically and applies in the
settings. Subscribe to the channel and
support me on the Patreon. The link in
the description.

---

*변환 완료: 2026-05-04 | 원본 무손실 전문 | AX 표준 적용*
*다음 파트: 없음 (최종)*


---
## 🔗 관련 시너지 지식
- **핵심 하네스**: [Global_Scenario_Harness_21.md](../../Global_Scenario_Harness_21.md)
- **클러스터**: Creative Production