{{$args := parseArgs 1 "To heal someone, type `-heal @<user>`" 
(carg "member" "member to heal")}}

{{$member := $args.Get 0}}
{{if (hasRoleID /* Replace with your Virus Role ID */)}} {{/* if user has the role Virus */}}
  {{if eq .User.ID $member.User.ID}} {{/* if target is self */}}
    A Virus cannot heal itself.
  {{else}}
    You cannot heal people as a Virus :microbe:.
  {{end}}
{{else}}
  {{if (targetHasRoleID $member.User.ID /* Replace with your Infected Role ID */)}} {{/* if target is infected */}}
    {{if eq .User.ID $member.User.ID}} {{/* if target is self */}}
      Sorry, you cannot heal yourself. Get a friend to help you! {{/* prevents self-healing */}}
    {{else}}
      {{takeRoleID $member.User.ID /* Replace with your Infected Role ID */}} {{/* removes infection from the target user */}}
      {{.User.Mention}} has healed {{$member.User.Mention}} from the infection!
    {{end}}
  {{else}}
    {{$member.User.Username}} is well and healthy!
  {{end}}
{{end}}