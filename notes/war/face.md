```sql
select n.name npc_name, n.level npc_level, s2.zone, n.hp npc_hp, CONCAT(n.mindmg, "-", n.maxdmg) npc_damage, i.name item_name, i.id as item_id, lde.chance, i.ac as item_ac, i.hp as item_hp
FROM items i
inner join lootdrop_entries lde on lde.item_id = i.id
inner join loottable_entries lte on lte.lootdrop_id = lde.lootdrop_id
inner join npc_types n on n.loottable_id = lte.loottable_id
inner join spawnentry se on se.npcID = n.id
inner join spawn2 s2 on s2.spawngroupID = se.spawngroupID
where
i.slots & 8 = 8
AND s2.zone IN (SELECT LOWER(short_name) from zone where expansion = 0)
and lde.chance > 0
AND i.classes & 1 = 1
AND i.ac > 1
GROUP BY item_name order by i.ac DESC;
```

#Water_Elemental_Overseer	61	paw	29000	135-550	"Mask of the Defeated "	32989	87.5	16	90
Cazic_Thule	70	fearplane	451000	223-603	"Eye of Cazic-Thule "	1547	10	15	50
#a_Tae_Ew_convert	56	cazicthule	14750	79-184	"Frightforged War Mask "	27855	15	12	35
Vorshar_the_Despised	61	neriakb	32000	85-185	"Impenetrable Mask of Al`Kabor "	24039	10	10	40
a_phantom_chest	57	hateplaneb	15750	84-370	"Triumphant Mask "	1539	0.347	10	0
#Mistress_of_Malevolence	60	hateplaneb	23000	74-368	"Bloody Mask of Fury "	62007	10	10	35
a_grotesque_rat	63	hateplaneb	29000	134-533	"Darkwar Mask "	2190	70	10	0
#an_enraged_Amygdalan	56	cazicthule	40000	78-230	"Amygdalan's Chain Mask "	27867	70	8	85
a_skeletal_commander	39	tox	1911	16-82	"Cold Visor of the Warlord "	69300	2	7	20
a_lifestealer_mosquito	53	cazicthule	11750	61-175	"Rubicite Mask "	4162	0.335	7	0
a_grim_abhorrent	58	hateplaneb	16750	84-370	"Darkbrood Mask "	1544	0.641	7	0
#War_Priestess_T`zan	59	kithicor	17000	62-181	"Teir`Dal Adamantite Visor "	21168	2	6	0
#Lord_Pickclaw	32	runnyeye	1644	13-68	"Grotesque Alloy Mask "	3618	14	6	0
#Dreadfang	58	cazicthule	20750	125-258	"Dreadfang Mask "	27946	35	6	50
a_virulent_mosquito	58	cazicthule	24750	85-220	"Mosquito Hide Mask "	27879	28.05	6	0
an_ice_goblin	16	permafrost	416	1-36	"Bronze Mask "	4202	0.056	5	0
an_ogre_guard	20	qey2hh1	600	1-44	"Large Bronze Mask "	4226	2.25	5	0
Warlord_Skarlon	53	soldungb	19138	77-213	"Polished Mithril Mask "	4505	20	5	0
a_ghastish_ancille	33	mistmoore	1419	13-70	"Small Bronze Mask "	4214	0.273	5	0
a_goblin_archeologist	27	permafrost	1200	12-64	"Dented Brass Mask "	3314	35	4	0
an_imp_familiar	29	mistmoore	1131	11-62	"Bloodstone Eyepatch "	1410	50	4	0
an_orc_warrior	10	kithicor	200	1-24	"Pickbringers Chainmail Faceguard "	6107	4.25	3	0
an_ice_goblin	16	permafrost	416	1-36	"Iron Visor "	3102	0.244	3	0
a_guard	36	highpass	1656	43-220	"Iron Mask "	3002	3.5	3	0
Watchman_Grep	35	steamfont	2722	14-74	"Small Iron Visor "	3114	1.25	3	0
an_ogre_shaman	27	qey2hh1	999	1-58	"Sheer Bone Mask "	10101	3.226	3	0
centaur_courser	25	southkarana	840	1-54	"Large Iron Visor "	3126	0.755	3	0
a_Tesch_Mas_Gnoll	25	southkarana	980	1-54	"Split Paw Hide Mask "	2282	1	3	0
a_Silverflank_guardian	30	feerrott	1200	12-64	"Lizardskin Tribal Mask "	2330	6.25	3	0
Squallslither	14	cauldron	336	1-32	"Serpentskin Eyepatch "	1317	25	3	0
#Lord_Pickclaw	32	runnyeye	1644	13-68	"Grotesque Mask "	3041	14	3	0
Lord_Gongo	20	kerraridge	600	1-44	"Gorilla Hide Mask "	2373	75	3	0
BoneCracker	24	najena	816	1-52	"Leering Mask "	2311	6.25	3	0
orc_legionnaire	12	crushbone	264	1-28	"Raw-hide Mask "	2138	1.236	2	0
orc_oracle	13	crushbone	299	1-30	"Small Raw-hide Mask "	2150	0.417	2	0
Lord_Darish	14	crushbone	336	1-32	"Large Raw-hide Mask "	2162	0.759	2	0
a_kobold_hunter	7	innothule	125	1-18	"Woven Mask "	1122	0.431	2	0
inferno_goblin_torturer	28	soldunga	1064	11-60	"Turquoise Eyepatch "	1411	100	2	0
a_gnoll_shaman	15	lakerathe	375	1-34	"Snakeskin Mask "	2307	6.78	2	0
Ekeros	18	najena	504	1-40	"Leather Mask "	2002	2.353	2	0
a_froglok_ton_shaman	19	guktop	551	1-42	"Mesh Mask "	1142	1.25	2	0
a_froglok_idealist	20	guktop	600	1-44	"Moss Mask "	1400	30	2	0
a_minotaur_elder	32	gukbottom	2000	13-78	"Carved Ivory Mask "	10144	29.167	2	10