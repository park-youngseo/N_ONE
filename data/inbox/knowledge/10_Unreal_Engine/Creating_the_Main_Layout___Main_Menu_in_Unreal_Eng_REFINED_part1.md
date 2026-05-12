# Creating the Main Layout | Main Menu in Unreal Engine 5.5.4 (Part 1)
## [메타데이터]

| 항목 | 내용 |
|------|------|
| **채널** | Unreal Fusion |
| **영상 ID** | Rh7FGqFj2Cs |
| **영상 URL** | https://www.youtube.com/watch?v=Rh7FGqFj2Cs |
| **카테고리** | 10_Unreal_Engine |
| **파트** | 1 / 2 |
| **변환일** | 2026-05-04 |

---

## [원본 자막 전문 — 무손실 (Part 1)]

> **[AX 하네스 준수]** 요약 및 압축 절대 금지. 원본 자막 텍스트 전문 보존.

First thing first, uh this first part of
the main menu
series, I want to create the main layout
of our main menu. Okay, first of all, we
have to open the engine and go to game.
Select third person as template and uh
choose the name of your
project menu
course. Hit
create. Okay, now it's open it. Let's go
and create a new folder. Name this one
main
new change the color
folder let's create another
folders
widgets
images.
This
and change the color of those
folders. Choose any color.
Okay, now let's go to the widget and
create go to the user interface and
create new
widget main menu.
equipment
hook. Okay. In other tutorials, we have
this uh old users, old YouTubers or the
person who creates courses. I'll use
canvas
panel but for optimization purposes
uh we have to avoid to use canvas
panel we have to use overlays instead
of overlays.
Drag this overlay right here. So this
one and this overlay the main
overlay. Then after that we have to
add an other overlay. Before I add that
this other we have to add a widget
switcher. We use the switch switcher uh
by index elements
uh to to switch between menus and
others. For example, between level
select and menu and setting and the main
menu and the counter of these widgets is
become from zero as indexed
to to your main uh layout. So you have
five layouts five, six and others. Okay,
let's rename this
switcher main
menu
which may trans widget feature.
We have to create another
overlap. This is the first
one. Okay. This one main menu
buttons
overlay. Okay. Select this
overlay like
this and make an alignment like this.
Okay, return the widget switcher like
this. I have in this case slots for the
first let's add those like this switch.
Now before we
creating button is we have to add a
background
image. I want to tell you about the free
website to add background images,
sounds, effect, and music. Okay, let's
go to the Google
Chrome and let's search for
Pixab. This website is allows you to
have
effect
English and sounds, free videos and
music and free
images. Okay, for us we have for to do
that we have to create an account in
this website and you can download the
images for
free. For us we have
to add background images.
There are any background what you want.
Okay, for me I want to choose cyberpunk
[Music]
background. Choose any background uh for
this
tutorial. Let's choose this
background. I think it's very
clean. it download and change the
quality as high just as
possible. Okay, let's return right
here. For now, it's
good. I want to download others in the
next two. Okay,
close open the folder.
and choose my images
menu. I drag this
background. Okay. Right here and select
image. Drag this image in this overlay.
Once it's selected, change this image to
bgs. Once it's selected, make it in the
main
layout and select the
brush and choose from
the the background image.
Now return right here and change it.
Make wider make it like
this. Now let's add our
buttons. Okay. Let's choose another
overlay to cover the
buttons. Okay. But
Okay, make it like this for this
moment. And now let's create a vertical
box.
And this vertical box we have to create
buttons. And this button we have to
create
text. Okay. Select the vertical box like
this and drag it right here. The align
vertical line it in the
center. Okay. Now, and let's make a
padding on
this cover
overlay 10
pixel. And for the
vertex, select the cover overlay like
this.
And let's add
finish. Okay. Make the image right here.
And where is the image? Right here.
Okay. Select the image and set it like
this. Once the image is selected, we run
to change the name of
[Music]
cover
overlay image.
Now select the color of this image as a
black and let's go to the opacity or
alpha make it like
this. And now this first button we have
to make it as starter me
start change from here or
here start
game once it's changed right
here. Now the text is
starting. Now let's modify the let's
change the background of the button. We
have to go to the tint right
now. Okay. Let's change it
to red.
And now it's draw as rounded
focus and go to the outline and change
little more. For example,
50. Okay, look at this
one. And the final
one is
50. In this case, okay, I don't want
this. I want this one to be 50.
Now it's
good
side. Okay. What do you want otherwise?
We have to go
here. Select the
brush. Like the
bottom. Now when one it's
hovered the t it changes to
orange.
Okay. What
else? When it's pressed it changes to
green or you choose change the color as
you need as you want.
Okay. And now it's fine for the first
button. Now let's select this button.
Ctrl D to
duplicate. Duplicate it four
time. Start game level selection.
sitting uh five
times.
Okay. Now select the second
one, third one, fourth one and last one
and return here to the padding from this
one. Change the top to
20. Now it's better. This looks better.
And now let's select second button and
change the
text this one
11.
Select
setting.
What? Okay, let's change it from this
second button. Select it. So, we go
right
here and change this
one. Like
it
but select this button and change this
for settings
button for the credit.
Edit
button. Final one
is okay. Now it's fine. It remains one
thing for this
tutorial. And we
ended file save part.
Okay, before we finish the tutorial, we
have to go
here and go here to the main
menu and create new
level. Main new level.
Select this level same as
selected. Go right here and on this one
open
level as an event being play
widget. Okay. Where is my widget? Me
preprint from the return value add to
view
port view
port. Okay. Save this. And return here
to my menu. Menu widget. Go to the
graph from the event
construct. Let's go here. Get player
controller. Once it selects it
most
[Music]
person is a
true for my main menu we want to show
the UI only.
Okay, let type UI
only and drag the player control like
this. Fox widget this one make it
set in mouse lock
button lock
always. Okay, make it like this.
Now let's go here and go to the content
person. For me it's a third person for
you. You can use another character
custom
character. Go down right here.
Event begin
play. Get player controller.
Show mouse
cursor for the
menu is true. When I start playing it,
make it to
false. I want to set the game only in
this type.
that input mode
game
on. Okay, hit compile and save it and
let's
see. Okay, it's problem from this player
control my player zone. Okay, to fix
this problem, we have to go
here from the main menu and click
new game mode P or the main
menu. Main
menu game mode.
Okay, once it's
selected, but if I go here to the
default power
plus and change it to my main
menu. Okay, change it to none. It's
fine.
After changing to none, go to the
setting project
setting. Head back in this place where
game
[Music]
mode. Change this one to my main menu.
Left. Second one main menu
and for the default mode here. Change it
to my main
menu. I think the problem is fixed.
Okay, let's
see. Okay,
fine. Our first part is of my main menu
is
done. Okay, thank you guys for watching.
Support me on the Pat videos like
this and subscribe for the channel
to make another tutorial to help me lot.

---

*변환 완료: 2026-05-04 | 원본 무손실 전문 | AX 표준 적용*
*다음 파트: part2*


---
## 🔗 관련 시너지 지식
- **핵심 하네스**: [Global_Scenario_Harness_21.md](../../Global_Scenario_Harness_21.md)
- **클러스터**: Creative Production
