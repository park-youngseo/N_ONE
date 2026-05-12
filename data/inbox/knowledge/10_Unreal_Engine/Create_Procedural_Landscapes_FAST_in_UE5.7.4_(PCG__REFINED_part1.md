# Create Procedural Landscapes FAST in UE5.7.4 (PCG Tutorial)
## [메타데이터]

| 항목 | 내용 |
|------|------|
| **채널** | Unreal Fusion |
| **영상 ID** | VxcP5guTU7I |
| **영상 URL** | https://www.youtube.com/watch?v=VxcP5guTU7I |
| **카테고리** | 10_Unreal_Engine |
| **파트** | 1 / 2 |
| **변환일** | 2026-05-04 |

---

## [원본 자막 전문 — 무손실 (Part 1)]

> **[AX 하네스 준수]** 요약 및 압축 절대 금지. 원본 자막 텍스트 전문 보존.

Hey guys, in this tutorial, I'll be
showing you how to create a
PCG forest like this. How to add trees,
foliage, grasses,
sand, water, and water, and fix some
general number of this PCG cap. Okay,
let's go.
Okay, everything first open the editor.
A new project.
Games, third-person master template.
Project name, whatever you like, you
know.
And hit create.
Now, let's go ahead to file, create new
level,
create basic level, and create.
Different than escape, and go here.
Select mode, landscape mode.
And hit create.
Return to select mode, and go here to
file,
save current level,
and hit save.
And now, let's go here
to plugins, edit plugins, and search for
PCG content generation.
PCG, and enable.
To enable this PCG, you need to restart
the engine.
Let's restart.
Okay, fine.
Go to here to content browser, new
level.
Okay, this is my level. Let's go right
here, select PCG, go PCG graph, create
empty graph.
I'll rename it to
PCG
C
graph.
Okay, fine.
Make this one like texture like this,
and go here.
And let's add my PCG graph in this
corner.
Okay, before I start working, we have to
go here,
content browser, content file, and
download some assets.
Search here for grass.
A more field tool lies on this software
product list.
Okay, let's download this.
Okay, accept.
Second one, all of the grass.
Add to my project like
Okay, this one.
And finally, this one.
Okay.
Return like this.
Close this.
And search
for
model in
Okay, this one.
European beech.
Download this asset, add to my project.
Download to the closest version for me,
so I this one, and hit add to my
projects. Take some time to download
this asset.
Okay, fine. Downloading is finished.
Let's return right here.
This is my European beech.
Go right here,
and go to static mesh in the blueprint.
Select all of them, control Q.
Or control play.
And enable the night. As you can see,
and the night is enabled by default.
Now,
next thing I want to do
is to go here,
and get the
landscape data.
Okay, I want to get this landscape
inside this view
data.
After that, we'll add surface sampler
surface sampler node,
and
let's add static
mesh
spawner.
In the static mesh spawner, we'll add
our trees.
Okay, select this one, and go right
here. There is those mesh entries.
I want to
add on
Okay, how many type of trees I wanted to
add?
Eight types of trees. Okay, this array
8 plus buttons.
And go change the weight, for example,
this one. Two,
one,
three,
two,
one.
And change it as you want.
And after that,
go here, search for static.
When you go here, and search for static,
we have to add or drag those meshes one
by one on this place.
Okay, let's see.
Okay, before we do that,
we have to go here, and add this to back
button, and let's see what's happening
here.
This
on this point, this is the place of our
trees.
And this uh
inside this cube.
Uh okay,
let's minimize this.
Before we add this,
select this one, and go here to sting.
Okay,
0.0
five.
0.0
five.
And let's see.
Okay, I think this is good.
Now, let's return right here, and
search. As you can see, this way static,
and continue our
work.
Okay, let's choose this one, first one.
Second one, let's choose this.
This one, third one.
Okay, this one.
Choose this.
And choose the
uh this.
Okay.
Should I remove this one?
Okay, let's choose, for example, this
one.
Okay, that's all the
save, and let's see what's happening
right here.
Okay, preparing texture.
Okay, once it's prepared, before we do
that, we have to go here,
and disable
our tree pack.
I can see
the first trees is appears in some time
to run the all those uh
uh trees.
I see them.
Okay, this is our trees. It's not a lot.
Okay, let's reduce this one to 0.01.
Okay.
Is 0.0 0.0 good?
Maybe 0.08.
Okay, let's add our third-person
character right here.
Okay, where is my
basic Let's add start play on this
place, and let's see
what's happening right here.
Play.
Okay, as you can see, this is my
third-person character, as you can see.
And this is my trees.
Okay, it's heavy for the CPU.
There is a lot of
uh trees with a lot of meshes,
and those meshes are heavy
for our
CPU.
Okay, let's return right here,
and go over here, and add a density
filter.
Density filter the output in this place.
Unplug this one, and go here
to our density.
Uh this is lower bound, make it like
this. Upper bound, make [clears throat]
it like this. And let's see.
Okay.
I think it's good right now.
>> Okay, and all the trees and looseness.
Then one thing.
Okay, now let's add the material for our
forest ground.
Let's disable the static mesh. Go to the
fog.
Before we do that, we have to save our
work.
Okay, let's see where the dissolve
by
grass.
Okay, set price for free.
And let's see this dried grass.
I want the forest.
Okay, I think this one
is a good.
Okay, let's use this.
Uh if you have good uh GPU, okay, let's
use high quality.
Okay, hit accept.
Okay, fine. Fine is done. Loading.
Let's select our landscape from both
down right here. And add this material
in this place. How to undo that?
We have to go here, fog, but we just
can't
the surfaces. Right grass uh
material.
And drag this one in this place.
Okay.
Okay, let's change it. Let's reduce the
tiling and add more normal. Go here to
global
tiling.
And I can change the X offset
and the Y offset.
Basic color.
Saturation.
Roughness.
Contrast.
And then water.
Let's convert tone.
Okay, for the specular
like this. So, um
And finally, the normal.
And let's see.
Okay.
Let's add mesh to present character.
Okay, now what I want to do
is to go again.
Let's shift
shift add new.
And then I add some shift in.
I'm sure.
Delete this.
Go to front view.
Okay, I think my forest now it's looks
uh
like good. Let's see.
Okay, now let's make some deformation.
In our landscape, before we do that, we
have to hide
our PCG graph.
And complete.
Go here. Let's hide it.
And return here.
Save all.
Landscape.
We make the little
brush effect.
Okay. And let's see. Return light here.
Let's Let's remove them. And let's see.
Okay. Listen, there is problem of
rotation of trees.
Let's fix the rotation.
Let's return it in this place.
Okay, let's return it to PCG graph.
And add another node the transform
points.
And plug this.
And then plug it in this place.
Look at this node and go here to
rotation. Select this absolute rotation

---

*변환 완료: 2026-05-04 | 원본 무손실 전문 | AX 표준 적용*
*다음 파트: part2*


---
## 🔗 관련 시너지 지식
- **핵심 하네스**: [Global_Scenario_Harness_21.md](../../Global_Scenario_Harness_21.md)
- **클러스터**: Creative Production
