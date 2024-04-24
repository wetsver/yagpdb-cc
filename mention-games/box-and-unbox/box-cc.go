{{$args := parseArgs 1 "To box someone, type `-box @<user>`" 
(carg "member" "member to box")}}
 
{{$member := $args.Get 0}}
{{if eq .User.ID $member.User.ID}}
  You cannot box yourself. Get someone to tape your box for you!
{{else if (hasRoleID /* Replace with your Boxed Role ID */)}} {{/* if user is already boxed */}}
  You are stuck in a box! :package:
{{else if (targetHasRoleID $member.User.ID /* Replace with your Boxed Role ID */)}}  {{/* if target is boxed */}}
  Sorry, {{$member.User.Username}} is stuck in another box :smirk_cat:
{{else if eq $member.User.ID 204255221017214977}} {{/* if target is YAGPDB */}}
  Noob :baby: {{/* Avoid infecting the bot*/}}
{{else}}
  {{giveRoleID $member.User.ID /* Replace with your Boxed Role ID */}} {{/* assigns boxed role to user */}}
  :smirk_cat: {{.User.Mention}} has used Box no Jutsu on {{$member.User.Mention}}!
{{end}}