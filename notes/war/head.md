```sql
select n.name npc_name, n.level npc_level, s2.zone, n.hp npc_hp, CONCAT(n.mindmg, "-", n.maxdmg) npc_damage, i.name item_name, i.id as item_id, lde.chance, i.ac as item_ac, i.hp as item_hp
FROM items i
inner join lootdrop_entries lde on lde.item_id = i.id
inner join loottable_entries lte on lte.lootdrop_id = lde.lootdrop_id
inner join npc_types n on n.loottable_id = lte.loottable_id
inner join spawnentry se on se.npcID = n.id
inner join spawn2 s2 on s2.spawngroupID = se.spawngroupID
where
i.slots & 4 = 4
AND s2.zone IN (SELECT LOWER(short_name) from zone where expansion = 0)
and lde.chance > 0
AND i.classes & 1 = 1
AND i.ac > 1
GROUP BY item_name order by i.ac DESC;
```

Gizula	65	Everfrost	36300	320-800	"Malevolent Helmet of Suffering "	44113	25	35	60
#an_ancient_construct	55	hole	13750	78-219	"Rock Helm of Defense "	62291	33	25	35
#a_gelatinous_mass	55	cazicthule	13750	78-219	"Tae Ew War Helm "	27844	50	24	25
a_kiraikuei	50	hateplane	11148	60-204	"Indicolite Helm "	4911	2.25	20	0
a_frightful_chest	60	hateplaneb	18000	93-370	"Crimson Helm of Horror "	62009	20	20	50
a_nightmare	51	fearplane	14285	48-143	"Legionnaire Scale Helm "	5174	2.639	19	0
#War_Priestess_T`zan	59	kithicor	17000	62-181	"Teir`Dal Adamantite Helm "	21167	2	13	0
Cauldronbubble	49	kedge	10128	55-133	"Hammerhead Helm "	4620	30	13	0
King_Tranix	52	soldungb	16250	76-210	"Crown of King Tranix "	4504	75	13	0
the_ghoul_lord	47	gukbottom	10863	50-126	"Skull-shaped Barbute "	4301	75	13	35
the_froglok_king	47	gukbottom	7250	73-149	"Crown of the Froglok Kings "	10142	75	12	0
a_lifestealer_mosquito	53	cazicthule	11750	61-175	"Rubicite Helm "	4161	0.223	12	0
a_Darkpaw_curate	42	jaggedpine	3750	34-130	"Sabertooth Forged Helm "	8037	8.5	12	0
an_ice_goblin	16	permafrost	416	1-36	"Bronze Helm "	4201	0.077	10	0
Noble_Dojorn	60	airplane	32000	98-326	"Golden Efreeti Turban "	4341	4	10	0
an_ogre_guard	20	qey2hh1	600	1-44	"Large Bronze Helm "	4225	2.5	10	0
Guard_Dunn	45	butcher	4277	36-139	"Small Bronze Helm "	4213	25	10	0
an_avenging_caitiff	30	mistmoore	1200	12-64	"Crested Helm "	4300	50	10	0
Hatar	20	oasis	600	1-44	"Fine Silk Turban "	1405	12.5	10	0
a_frightful_chest	60	hateplaneb	18000	93-370	"Crown of Hatred "	2191	80	10	0
Urwenae_the_Cold	55	hole	13750	74-169	"Paineel Platemail Helm "	4381	0.4	10	0
#Queen_Nhyalia	53	jaggedpine	20625	77-213	"Enchanted Dew Woven Circlet "	8975	50	10	20
an_orc_warrior	10	kithicor	200	1-24	"Pickbringers Chainmail Coif "	6010	4.25	8	0
a_windrider_drake	60	airplane	19000	237-655	"Sphinx Crown "	20871	13.223	8	8
#Lord_Pickclaw	32	runnyeye	1644	13-68	"Blackened Alloy Coif "	3600	10	8	0
a_deepwater_goblin	35	lakerathe	1575	14-74	"Opoline Helm "	4099	8.75	8	0
#Coercer_Q`ioul	58	kithicor	18000	61-178	"Teir`Dal Chain Coif "	3283	6.931	7	0
an_ice_goblin	16	permafrost	416	1-36	"Ringmail Coif "	3101	0.24	7	0
a_guard	36	highpass	1656	43-220	"Chain Coif "	3001	3.5	7	0
a_gnoll_soothsayer	16	highpass	416	1-36	"Large Ringmail Coif "	3125	12.5	7	0
Den_Ironblend	35	butcher	2200	14-74	"Small Ringmail Coif "	3113	12.5	7	0
a_goblin_magician	19	najena	551	1-42	"Stone Circlet "	10122	3.75	6	0
Chief_Goonda	34	qey2hh1	2100	14-77	"Turtleshell Helm "	10100	3.75	5	0
a_ghoul_executioner	35	gukbottom	2575	14-74	"Executioners Hood "	2468	75	5	0
the_ghoul_lord	47	gukbottom	10863	50-126	"Scalp of the Ghoul Lord "	26997	100	5	0
#Tzitzi_the_Crazed	52	hole	10750	45-144	"Earthen Crown of the Underfoot "	8917	95	5	10
Quid_Rilstone	34	rathemtn	1644	32-76	"Siryn Hair Hood "	1501	40	5	0
orc_legionnaire	12	crushbone	264	1-28	"Raw-hide Skullcap "	2137	1.09	4	0
orc_oracle	13	crushbone	299	1-30	"Small Raw-hide Skullcap "	2149	1.835	4	0
Lord_Darish	14	crushbone	336	1-32	"Large Raw-hide Skullcap "	2161	0.506	4	0
a_goblin_preacher	24	permafrost	816	1-52	"Runed Circlet "	3400	60	4	0
a_gelatinous_cube	5	qcat	75	1-14	"Cat Skull Cap "	12372	5	4	0
Leatherfoot_Scout	1	nektulos	11	1-6	"Leatherfoot Raider Skullcap "	13113	25	4	0
a_froglok_tad	1	feerrott	11	1-6	"Mesh Helm "	1141	1.5	4	0
Shazda_Asad	23	kerraridge	759	1-50	"Leather Skullcap "	2001	12.5	4	0
Nillipuss	15	rivervale	700	1-34	"Small Leather Skullcap "	2013	1.125	4	0
a_froglok_forager	6	innothule	100	1-16	"Woven Cap "	1121	0.847	3	0
a_gnoll	16	lakerathe	416	1-36	"Ruined Gnoll Chain Coif "	20109	2	3	0
Spinflint	12	beholder	267	1-28	"Whirling Crystal "	10205	8	3	10
Mroon	20	southkarana	600	1-44	"Topknot Headband "	2320	12.5	3	0
orc_slaver	5	crushbone	75	1-14	"Cloth Cap "	1001	1.342	2	0
orc_centurion	9	crushbone	171	1-22	"Small Cloth Cap "	1013	1	2	0
a_froglok	4	qcat	56	1-12	"Netted Cap "	1101	2.25	2	0
a_snow_orc_trooper	7	Everfrost	155	4-18	"Large Cloth Cap "	1025	1.203	2	0
an_ogre_shaman	27	qey2hh1	999	1-58	"Batfang Headband "	10108	2.151	2	0
a_gnoll	3	qeytoqrg	39	1-10	"Ruined Gnoll Leather Skullcap "	20119	0.127	2	0
Lord_Bergurgle	40	lakerathe	3250	32-124	"Lord Bergurgle's Crown "	28044	100	2	0
reclusive_ghoul_magus	23	unrest	859	1-50	"Savant's Cap "	1302	80	2	0
a_fire_drake	9	lavastorm	171	5-25	"Ball of Primeval Lava "	51889	0.251	2	0