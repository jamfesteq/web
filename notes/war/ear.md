select n.name npc_name, n.level npc_level, s2.zone, n.hp npc_hp, CONCAT(n.mindmg, "-", n.maxdmg) npc_damage, i.name item_name, i.id as item_id, lde.chance, i.ac as item_ac, i.hp as item_hp
FROM items i
inner join lootdrop_entries lde on lde.item_id = i.id
inner join loottable_entries lte on lte.lootdrop_id = lde.lootdrop_id
inner join npc_types n on n.loottable_id = lte.loottable_id
inner join spawnentry se on se.npcID = n.id
inner join spawn2 s2 on s2.spawngroupID = se.spawngroupID
where
i.slots & 2 = 2
AND i.slots & 16 = 16
AND s2.zone IN (SELECT LOWER(short_name) from zone where expansion = 0)
and lde.chance > 0
AND i.classes & 1 = 1
AND i.ac > 1
GROUP BY item_name order by i.ac DESC;



#Guardian_of_the_Seal	70	hole	124000	300-900	"Purified Steel Earring "	62297	32	15	85
A_Legendary_Behemoth	72	steamfont	400000	400-1404	"Clawed Earthcrafter's Hoop "	22972	4	15	140
#Water_Elemental_Officer	61	paw	29000	135-555	"Defeated Gnoll's Earring "	32937	27.723	15	90
a_deathly_usher	30	mistmoore	1200	12-64	Silverwing Loop	79621	1.836	8	15
#a_diseased_mosquito	58	cazicthule	23750	81-228	"Mosquito Leg Earring "	27820	50	8	65
A_Legendary_Behemoth	72	steamfont	400000	400-1404	"Whispering Air Bangle "	2630	4	8	140
#a_toxic_jungle_hunter	56	cazicthule	14750	79-222	"Spiderfang Earring "	27822	50	7	35
#a_Thul_Tae_Ew_torturer	60	cazicthule	25000	99-270	"Earring of Pain "	27941	50	6	35
a_timber_griffon	39	jaggedpine	1911	16-82	"Griffon Fang Earring "	8291	2.5	6	10
#a_Thul_Tae_Ew_ritualist	60	cazicthule	25000	99-210	"Lizard Tooth Earring "	27942	25	5	75
#Assassin_Z`Jrix	61	hateplaneb	21000	134-533	"Earring of Eternal Alacrity "	62026	20	5	30
a_vicar_of_hate	58	hateplaneb	16750	84-370	"Earring of Bashing "	1538	0.452	5	0
an_Agent_of_Innoruuk	54	hateplane	11139	78-216	"Silver Earring of the Chosen "	62041	25	4	65
a_spirit_of_the_old_city	51	hole	9000	53-167	"Shiny Gold Earring "	62283	0.015	4	35
a_skeletal_commander	39	tox	1911	16-82	"Dark Hoop of the Necromancer "	69303	2	3	15
orc_taskmaster	14	crushbone	336	1-32	"Brass Earring "	10351	35	2	0
Singe	28	soldunga	1064	11-60	"Obsidian Bead Hoop "	10339	75	2	0
Princess_Cherista	35	mistmoore	1575	14-74	"Diamondine Earring "	10165	34	2	12
an_undead_barkeep	25	unrest	875	1-54	"Opalline Earring "	10407	35	2	0
#a_Thul_Tae_Ew_adept	60	cazicthule	32000	99-270	"Lizard Bone Earring "	27948	50	2	70





