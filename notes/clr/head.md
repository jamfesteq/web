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