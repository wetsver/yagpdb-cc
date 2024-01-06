{{$args := parseArgs 1 "To infect someone, type `-infect @<user>`" 
(carg "member" "member to infect")}}

{{$member := $args.Get 0}}
{{if (hasRoleID /* Replace with your Virus Role ID */)}} {{/* if user has the role Virus */}}
  {{if (targetHasRoleID $member.User.ID /* Replace with your Virus Role ID */)}}  {{/* if target is a Virus */}}
    A Virus cannot be infected.
  {{else if (targetHasRoleID $member.User.ID /* Replace with your Infected Role ID */)}} {{/* if target is infected */}}
    Sorry, {{$member.User.Username}} is already infected!
  {{else if eq $member.User.ID 204255221017214977}} {{/* if target is YAGPDB */}}
    Noob :baby: {{/* Avoid infecting the bot*/}}
  {{else}}
    {{giveRoleID $member.User.ID /* Replace with your Infected Role ID */}} {{/* assigns infected role to user */}}
    {{.User.Mention}} has infected {{$member.User.Mention}}!
  {{end}}
{{else}}
    You must be a Virus to infect other users.
{{end}}