```sql
select n.name npc_name, i.name item_name, i.id as item_id
FROM items i
inner join lootdrop_entries lde on lde.item_id = i.id
inner join loottable_entries lte on lte.lootdrop_id = lde.lootdrop_id
inner join npc_types n on n.loottable_id = lte.loottable_id
inner join spawnentry se on se.npcID = n.id
inner join spawn2 s2 on s2.spawngroupID = se.spawngroupID
where s2.zone IN (SELECT LOWER(short_name) from zone where expansion = 0) AND i.classes & 1 == 1 AND i.ac > 1 order by i.ac DESC;
```
Item Slot ID 	Item Slot Bitmask 	Item Slot Name
0 	1 	Charm
1 	2 	Ear 1
2 	4 	Head
3 	8 	Face
4 	16 	Ear 2
5 	32 	Neck
6 	64 	Shoulder
7 	128 	Arms
8 	256 	Back
9 	512 	Bracer 1
10 	1024 	Bracer 2
11 	2048 	Range
12 	4096 	Hands
13 	8192 	Primary
14 	16384 	Secondary
15 	32768 	Ring 1
16 	65536 	Ring 2
17 	131072 	Chest
18 	262144 	Legs
19 	524288 	Feet
20 	1048576 	Waist
21 	2097152 	Powersource
22 	4194304 	Ammo