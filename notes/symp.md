symp procs
https://retributioneq.com/allaclone/?a=spell&id=24437
```sql
select n.name npc_name, n.level npc_level, i.name item_name, lde.chance from items i
inner join lootdrop_entries lde on lde.item_id = i.id
inner join loottable_entries lte on lte.lootdrop_id = lde.lootdrop_id
inner join npc_types n on n.loottable_id = lte.loottable_id
where i.id in (select id-2000000 from items i where i.clickeffect = 24437 order by id asc) group by npc_name;
```
select id-2000000 from items i where i.clickeffect = 24437;

