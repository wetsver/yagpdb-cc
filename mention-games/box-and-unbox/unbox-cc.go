{{$args := parseArgs 1 "To rescue someone from the box, type `-unbox @<user>`" 
(carg "member" "member to unbox")}}
 
{{$member := $args.Get 0}}
{{if (hasRoleID /* Replace with your Boxed Role ID */)}} {{/* if user is stuck in a box */}}
  {{if eq .User.ID $member.User.ID}} {{/* if target is self */}}
    Someone is guarding your box. AAAHH!! :scream: :sob:
  {{else}}
    You are stuck in a box and cannot free others!
  {{end}}
{{else if (targetHasRoleID $member.User.ID /* Replace with your Boxed Role ID */)}} {{/* if target is boxed */}}
  {{takeRoleID $member.User.ID /* Replace with your Boxed Role ID */}} {{/* removes box from the target user */}}
  {{.User.Mention}} has rescued {{$member.User.Mention}} from the box! :tada:
{{else}}
  {{$member.User.Username}} is roaming around freely!
{{end}}