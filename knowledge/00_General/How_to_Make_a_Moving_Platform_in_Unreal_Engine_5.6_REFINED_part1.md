# How to Make a Moving Platform in Unreal Engine 5.6
## [메타데이터]

| 항목 | 내용 |
|------|------|
| **채널** | Unreal Fusion |
| **영상 ID** | r8arVgor2XA |
| **영상 URL** | https://www.youtube.com/watch?v=r8arVgor2XA |
| **카테고리** | 10_Unreal_Engine |
| **파트** | 1 / 1 |
| **변환일** | 2026-05-04 |

---

## [원본 자막 전문 — 무손실 (Part 1)]

> **[AX 하네스 준수]** 요약 및 압축 절대 금지. 원본 자막 텍스트 전문 보존.

Hi guys, in this tutorial I want to show
you how to make moving platforms like
this.
Okay, let's go.
Okay, first thing first, open the
editor. Go to game. Select the person at
template. Rename the project moving
platform.
Okay, go down here to content and
drawer. Create new actor blueprint.
Rename this actor
moving
platform
blueprint BPM.
Open it.
Now go down here to variable and create
a variable. The first one
is direction.
Set the type of this variable vector.
Second one selected and go here. Set the
speed
type of this variable is float. Now go
to the event graph compile. Save all.
Let's get the speed
here
and the direction here from this vector
formulas
normalize this vector.
After that we have to go here
multiply by the speed.
Okay, let's add and multiply the speed
by the world seconds get
[Music]
delta
seconds by the time from this you have
to multiply it.
Okay, after that we have to go here and
add MB
want to add a cube.
Let's make material for this cube.
Choose any material.
Okay, for now let's choose each one.
Okay, now let's get the cube right here.
And this get wet
transform.
Now get this cube.
Set the world
transform.
Drag the event tick right here and make
it here.
Now let's split this.
Split the other one.
Okay. From this let's add
the location to the location.
From that let's add to my new transform
and location
rotation value this one and scale value
this one. I think it's good right now.
Let's compile. Save all.
Go right here. We need a value for speed
100 by default.
Now let's set direction
here.
Duplicate this one. Ctrl C. Control V.
Now let's add D
of 3 seconds.
Flip plot
A and B.
First one is 10.
Second one is minus 10.
Now let's go here
and let's add my cube.
Make it right here.
Make it like that.
Okay. And let's see it works or not.
Okay, I think it's fine.
Now let's adjust the delay
to 6 seconds.
Let's see.
Okay, fine.
We can add the speed here. What do you
want?
Okay. Now, let's set the another moving
platform for the Y-axis. Okay. Duplicate
this.
Moving pl Y axis and this one Z-axis.
Okay. This one in this side and the
other one here.
Let's see the difference between them.
Okay, let's make it
this one right here.
Open the second one again.
Make the x-axis zero.
This one 20
minus 20.
Okay, let's search thread one
zero
and this one zero
300.
Select the cube and change the material.
This one is blue. Fine. Compile.
Select the cube and change this
material.
That's orange.
Okay. See?
And let's see it works or not.
Okay. X-axis and Zaxis Yaxis. Okay,
thank you guys for watching this all for
this tutorial. See you later in the next
tutorial.

---

*변환 완료: 2026-05-04 | 원본 무손실 전문 | AX 표준 적용*
*다음 파트: 없음 (최종)*


---
## 🔗 관련 시너지 지식
- **핵심 하네스**: [Global_Scenario_Harness_21.md](../../Global_Scenario_Harness_21.md)
- **클러스터**: Creative Production