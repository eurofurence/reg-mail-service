-- keeping these for reference, not currently in use
SET CHARACTER SET 'utf8mb4';

INSERT INTO `mail_templates` (`id`, `cid`, `lang`, `subject`, `data`, `created_at`, `updated_at`, `deleted_at`) VALUES
    ('16e44603-4abb-11ec-9cd3-3431c4db8789', 'roomreq_invite', 'de-DE', 'Gruppen Einladung', 'Liebe(r) {{ nickname }},\r\n\r\n* DU WURDEST EINGELADEN, EINER ZIMMERGRUPPE BEIZUTRETEN.\r\n\r\n{{ room_group_owner }} lädt dich dazu ein, in die Gruppe\r\n"{{ room_group_name }}" einzutreten.\r\n\r\nDu kannst ihn/sie unter {{ room_group_owner_email }} erreichen.\r\n\r\nWenn du die Einladung annehmen willst, musst du dich in das EF-\r\nAnmeldesystem einloggen und der Gruppe beitreten.\r\n\r\nSchöne Grüße,\r\n\r\n    Das EF Anmeldesystem.', '2021-11-21 12:06:34', '2021-11-21 12:06:34', NULL),
    ('c6bbd0c2-4aba-11ec-9cd3-3431c4db8789', 'roomreq_invite', 'en-US', 'Invited to Group', 'Dear {{ nickname }},\r\n\r\n* YOU HAVE BEEN INVITED TO BE SOMEONE\'S ROOMMATE!\r\n\r\n{{ room_group_owner }} invites you to join the group called\r\n"{{ room_group_name }}".\r\n\r\nYou can reach him/her at {{ room_group_owner_email }}.\r\n\r\nIf you wish to accept the invitation, you must log into the registration\r\nsystem and join the group.\r\n\r\nYours,\r\n\r\n    the EF registration system.', '2021-11-21 12:04:20', '2021-11-21 13:45:00', NULL),
	('454ccb05-4abd-11ec-9cd3-3431c4db8789', 'roomreq_leave', 'en-US', 'Someone left the Group', 'Dear {{ nickname }},\r\n\r\n* THIS IS AN AUTOMATED MAIL ABOUT YOUR ROOM SHARING GROUP\r\n  PLEASE READ IT, AND KEEP IT!\r\n\r\nWe would like to inform you that\r\n\r\n{{ room_group_member }}\r\n\r\nhas left your group.\r\n\r\n* HOW TO MANAGE YOUR GROUP\r\n\r\nYou can always log into the system and view your group\'s status using your\r\nlogin and password. You can remove anyone except yourself, and you can\r\ndelete the entire group.\r\n\r\nYou CANNOT actively add members to your group. They have to join by\r\nthemselves. They will need to know the name of your group to be able to\r\njoin. You can tell them the name of your group by any means you want, or\r\nuse the system\'s "invite" function.\r\n\r\n* IMPORTANT\r\n\r\nIf you wish to hand responsibility for your group to someone else,\r\nplease write to ef@eurofurence.org.\r\n\r\nYours,\r\n\r\n    the EF registration system.', '2021-11-21 12:22:11', '2021-11-21 13:45:58', NULL),
	('56b92c58-4aba-11ec-9cd3-3431c4db8789', 'roomreq_delete', 'en-US', 'The Group was Deleted', 'Dear {{ nickname }},\r\n\r\n* THIS IS AN AUTOMATED MAIL ABOUT YOUR ROOM REQUEST.\r\n\r\nThe group called\r\n\r\n"{{ room_group_name }}"\r\n\r\nhas been deleted. You are no longer in any group.\r\n\r\nThis is normally done by the creator of the group, or in very\r\nrare cases by the admins of the registration system.\r\n\r\n* WHAT CAN YOU DO NOW?\r\n\r\nYou can log into the registration system and either join a different group\r\n(you will have to know its name to do so) or create your own.\r\n\r\nYours,\r\n\r\n    the EF registration system.', '2021-11-21 12:01:12', '2021-11-21 13:45:50', NULL),
	('71a1fa6c-4abc-11ec-9cd3-3431c4db8789', 'roomreq_owner', 'de-DE', 'Gruppen Besitzer', 'Liebe(r) {{ nickname }},\r\n\r\n* DIES IST EINE AUTOMATISCHE BENACHRICHTIGUNG BETREFFS DEINER\r\n  ZIMMERGRUPPE. BITTE LESEN, UND NICHT LöSCHEN!\r\n\r\nDu bist ab sofort die neue verantwortliche Person für die Gruppe:\r\n\r\n"{{ room_group_name }}"\r\n\r\nDas bedeutet, dass du derjenige bist, der das Sagen darüber hat, wer der\r\nGruppe beitreten darf, und wer nicht.\r\n\r\nBitte beachte, deine Zimmergruppe ist nur ein Wunsch. Wir können nicht\r\ngarantieren, dass wir alle Furries deiner Gruppe wirklich in einem Raum\r\nunterbringen können. Wenn es machbar ist, werden Gruppen in Zimmern\r\npassender Größe untergebracht.\r\n\r\n* KURZANLEITUNG\r\n\r\nGebe den Namen deiner Gruppe jedem, den du gerne in deiner Gruppe hättest.\r\nSag ihnen, sie sollen sich über das Online-Anmeldesystem deiner Gruppe\r\nanschließen.\r\n\r\nAlternativ benutze die "invite" Funktion, während du im Anmeldesystem\r\neingeloggt bist.\r\n\r\n* DETAILLIERTE ANLEITUNG\r\n\r\nDu kannst NICHT aktiv Leute deiner Gruppe hinzufügen. Das müssen diese\r\nvon sich aus tun.\r\n\r\nSie brauchen dazu den Namen deiner Gruppe. Diesen kannst du ihnen auf\r\njedem beliebigen Wege mitteilen, oder die Einladungsfunktion im\r\nAnmeldesystem benutzen.\r\n\r\nNochmal, dein Gruppenname lautet:\r\n"{{ room_group_name }}"\r\n\r\nLeute, die noch auf der Warteliste stehen oder abgesagt haben, können\r\nkeiner Gruppe beitreten.\r\n\r\n* SO VERWALTEST DU DEINE GRUPPE\r\n\r\nDu kannst dich jederzeit in das System einloggen, und den Status deiner\r\nGruppe anschauen.\r\n\r\nDu kannst alle Mitglieder aus der Gruppe werfen außer dir selbst.\r\nAusserdem kannst du die gesamte Gruppe löschen.\r\n\r\nAls Verantwortlicher bekommst du jedesmal eine Mail, wenn sich an deiner\r\nGruppe etwas ändert, z.B. Neuzugänge oder Abmeldungen.\r\n\r\n* WENN DU HILFE BRAUCHST\r\n\r\nSchreibe an ef@eurofurence.org ... aber bitte mach das nur, wenn du\r\nwirklich nicht mehr weiterkommst. Wir erwarten viele Gäste, und wir machen\r\ndas auch nur in unserer Freizeit :)\r\n\r\n* WICHTIG\r\n\r\nWenn du die Verantwortung über deine Gruppe an jemand anderes abgeben\r\nwillst, bitte schreibe an ef@eurofurence.org.\r\n\r\nSchöne Grüße,\r\n\r\n    Das EF Anmeldesystem.', '2021-11-21 12:16:16', '2021-11-21 12:16:16', NULL),
	('75224306-4abd-11ec-9cd3-3431c4db8789', 'roomreq_leave', 'de-DE', 'Jemand hat die Gruppe verlassen', 'Liebe(r) {{ nickname }},\r\n\r\n* DIES IST EINE AUTOMATISCHE BENACHRICHTIGUNG BETREFFS DEINER\r\n  ZIMMERGRUPPE. BITTE LESEN, UND NICHT LÖSCHEN!\r\n\r\nWir möchten dich darüber informieren, dass\r\n\r\n{{ room_group_member }}\r\n\r\ndeine Zimmergruppe verlassen hat.\r\n\r\n* SO VERWALTEST DU DEINE GRUPPE\r\n\r\nDu kannst dich jederzeit in das System einloggen, und den Status deiner\r\nGruppe anschauen. Du kannst alle Mitglieder aus der Gruppe werfen,\r\nausser dir selbst. Ausserdem kannst du die gesamte Gruppe löschen.\r\n\r\nDu kannst NICHT aktiv Leute deiner Gruppe hinzufügen. Das müssen diese\r\nvon sich aus tun. Sie brauchen dazu den Namen deiner Gruppe. Diesen kannst\r\ndu ihnen auf jedem beliebigen Wege mitteilen, oder die "Invite"-Funktion im\r\nAnmeldesystem benutzen, um ihnen eine Einladung per E-Mail zuzuschicken.\r\n\r\n* WICHTIG\r\n\r\nWenn du die Verantwortung über deine Gruppe an jemand anderes abgeben\r\nwillst, dann schreibe bitte an ef@eurofurence.org.\r\n\r\nSchöne Grüße,\r\n\r\n    Das EF Anmeldesystem.', '2021-11-21 12:23:31', '2021-11-21 13:45:43', NULL),
	('7ba73b12-4abc-11ec-9cd3-3431c4db8789', 'roomreq_remove', 'en-US', 'Removed from Group', 'Dear {{ nickname }},\r\n\r\n* THIS IS AN AUTOMATED MAIL ABOUT YOUR ROOM REQUEST.\r\n\r\nYou have been removed from the group called\r\n\r\n"{{ room_group_name }}".\r\n\r\nThis is normally done by the creator of the group, or in very\r\nrare cases by the admins of the registration system.\r\n\r\nThe group was created by {{ room_group_owner }}, you can reach them at\r\n\r\n{{ room_group_owner_email }}\r\n\r\nif you should have questions as to the reasons for the removal.\r\n\r\nPlease do not simply add yourself again, but rather clear up the situation\r\nwith the group creator. Warning: Repeatedly adding yourself to groups whose\r\ncreators do not wish to have you in there may in extreme cases get you\r\nbanned from the convention for disruptive behaviour.\r\n\r\n* WHAT CAN YOU DO NOW?\r\n\r\nYou can log into the registration system and either join a different group\r\n(you will have to know its name to do so) or create your own.\r\n\r\nYours,\r\n\r\n    the EF registration system.', '2021-11-21 12:16:33', '2021-11-21 12:20:09', NULL),
	('8762815b-4abb-11ec-9cd3-3431c4db8789', 'roomreq_join', 'en-US', 'Someone joined the Group', 'Dear {{ nickname }},\r\n\r\n* THIS IS AN AUTOMATED MAIL ABOUT YOUR ROOM SHARING GROUP\r\n  PLEASE READ IT, AND KEEP IT!\r\n\r\nA furry named\r\n\r\n{{ room_group_member }}\r\n\r\nhas joined your group, and thus wants to be one of your roommates.\r\nYou can reach that person by writing to:\r\n\r\n{{ room_group_member_email }}\r\n\r\nIf you do not wish to have that person in your group, you can log into the\r\nEurofurence registration system and remove them. An automated note will be\r\nsent to that person if you do this.\r\n\r\n* HOW TO MANAGE YOUR GROUP\r\n\r\nYou can always log into the system and view your group\'s status using your\r\nlogin and password. You can remove anyone except yourself, and you can\r\ndelete the entire group.\r\n\r\nYou CANNOT actively add members to your group. They have to join by\r\nthemselves. They will need to know the name of your group to be able to\r\njoin. You can tell them the name of your group by any means you want, or\r\nuse the system\'s "invite" function.\r\n\r\n* IMPORTANT\r\n\r\nIf you wish to hand responsibility for your group to someone else,\r\nplease write to ef@eurofurence.org.\r\n\r\nYours,\r\n\r\n    the EF registration system.', '2021-11-21 12:09:43', '2021-11-21 13:45:15', NULL),
	('8a946b12-4aba-11ec-9cd3-3431c4db8789', 'roomreq_delete', 'de-DE', 'Die Gruppe wurde Gelöscht', 'Liebe(r) {{ nickname }},\r\n\r\n* DIES IST EINE AUTOMATISCHE BENACHRICHTIGUNG BETREFFS DEINER\r\n  ZIMMERGRUPPE.\r\n\r\nDu warst Mitglied bei der Zimmergruppe mit dem Namen\r\n\r\n"{{ room_group_name }}".\r\n\r\nDiese Gruppe wurde soeben gelöscht. Du bist nun in keiner Gruppe mehr.\r\n\r\nIn der Regel wird der Ansprechpartner der Gruppe die Löschung\r\nveranlasst haben. In seltenen Fällen tun das auch wir, aber in dem Fall\r\nbenachrichtigen wir den Ansprechpartner ebenfalls.\r\n\r\n* WAS NUN?\r\n\r\nDu kannst dich beim Anmeldesystem einloggen und entweder einer anderen\r\nGruppe beitreten (wenn du ihren Namen kennst) oder eine eigene anlegen.\r\n\r\nSchöne Grüße,\r\n\r\n    Das EF Anmeldesystem.', '2021-11-21 12:02:39', '2021-11-21 13:45:25', NULL),
	('c40d186b-4abb-11ec-9cd3-3431c4db8789', 'roomreq_join', 'de-DE', 'Gruppen Beitritt', 'Liebe(r) {{ nickname }},\r\n\r\n* DIES IST EINE AUTOMATISCHE BENACHRICHTIGUNG BETREFFS DEINER\r\n  ZIMMERGRUPPE. BITTE LESEN, UND NICHT LÖSCHEN!\r\n\r\n\r\nEin Furry mit dem Namen\r\n\r\n{{ room_group_member }}\r\n\r\nhat sich deiner Zimmergruppe angeschlossen. Du kannst ihn/sie unter\r\nder folgenden Adresse erreichen:\r\n\r\n{{ room_group_member_email }}\r\n\r\nWenn du diese Person nicht in deiner Gruppe haben möchtest, kannst\r\ndu dich beim EF-Anmeldesystem einloggen und ihn oder sie aus der Gruppe\r\nherauswerfen. Die Person bekommt automatisch eine EMail mit einer\r\nMitteilung.\r\n\r\n* SO VERWALTEST DU DEINE GRUPPE\r\n\r\nDu kannst dich jederzeit in das System einloggen, und den Status deiner\r\nGruppe anschauen. Du kannst alle Mitglieder aus der Gruppe werfen,\r\nausser dir selbst. Ausserdem kannst du die gesamte Gruppe löschen.\r\n\r\nDu kannst NICHT aktiv Leute deiner Gruppe hinzufügen. Das müssen diese\r\nvon sich aus tun. Sie brauchen dazu den Namen deiner Gruppe. Diesen kannst\r\ndu ihnen auf jedem beliebigen Wege mitteilen, oder die "Invite"-Funktion im\r\nAnmeldesystem benutzen, um ihnen eine Einladung per E-Mail zuzuschicken.\r\n\r\n* WICHTIG\r\n\r\nWenn du die Verantwortung über deine Gruppe an jemand anderes abgeben\r\nwillst, dann schreibe bitte an ef@eurofurence.org.\r\n\r\nSchöne Grüße,\r\n\r\n    Das EF Anmeldesystem.', '2021-11-21 12:11:25', '2021-11-21 12:23:50', NULL),
  	('f9709f26-4abc-11ec-9cd3-3431c4db8789', 'roomreq_remove', 'de-DE', 'Aus Gruppe entfernt', 'Liebe(r) {{ nickname }},\r\n\r\n* DIES IST EINE AUTOMATISCHE BENACHRICHTIGUNG BETREFFS DEINER\r\n  ZIMMERGRUPPE.\r\n\r\nDu warst Mitglied bei der Zimmergruppe mit dem Namen\r\n\r\n"{{ room_group_name }}".\r\n\r\nDu bist soeben ausgetragen worden!\r\n\r\nIn der Regel wird der Ansprechpartner der Gruppe die Löschung\r\nveranlasst haben. In seltenen Fällen tun das auch wir, aber in dem Fall\r\nbenachrichtigen wir den Ansprechpartner ebenfalls.\r\n\r\nDer Ansprechpartner war {{ room_group_owner }},\r\nDu erreichst ihn/sie unter {{ room_group_owner_email }},\r\nfalls du zu der Austragung eine Frage haben solltest.\r\n\r\nBitte sei so nett, und trage Dich nicht einfach sofort wieder ein, ohne mit\r\ndem Ansprechpartner geklärt zu haben, warum Du ausgetragen wurdest.\r\nWarnung: Wer sich wiederholt in Gruppen ohne Zustimmung der\r\nAnsprechpartner einträgt und uns damit ärger macht, kann in extremen\r\nFällen wegen Störung der Conplanung von der Teilnahme an Eurofurence\r\nausgeschlossen werden.\r\n\r\n* WAS NUN?\r\n\r\nDu kannst dich beim Anmeldesystem einloggen und entweder einer anderen\r\nGruppe beitreten (wenn du ihren Namen kennst) oder eine eigene anlegen.\r\n\r\nSchöne Grüße,\r\n\r\n    Das EF Anmeldesystem.\r\n', '2021-11-21 12:20:04', '2021-11-21 12:20:04', NULL),
	('fb729ef8-4abb-11ec-9cd3-3431c4db8789', 'roomreq_owner', 'en-US', 'Group Owner', 'Dear {{ nickname }},\r\n\r\n* THIS IS AN AUTOMATED MAIL ABOUT YOUR ROOM REQUEST.\r\n  PLEASE READ IT, AND KEEP IT!\r\n\r\nYou are now the new responsible contact person for the group called\r\n\r\n"{{ room_group_name }}"\r\n\r\nThat means you now have the ultimate say in who is allowed into it and who\r\nisn\'t.\r\n\r\nNote that your group is just a request. We cannot guarantee that we will be\r\nable to fit everyone in your group into one room. But we\'ll do our best.\r\nIf possible, groups will get rooms that match their size.\r\n\r\n\r\n* SHORT INSTRUCTIONS\r\n\r\nGive the name of your group to anyone you\'d like to join. Tell them to add\r\nthemselves using the Eurofurence online registration system.\r\n\r\nAlternatively, use the "invite" function while being logged into the online\r\nregistration website.\r\n\r\n* DETAILED INSTRUCTIONS\r\n\r\nYou CANNOT actively add members to your group. They have to join by\r\nthemselves.\r\n\r\nThey will need to know the name of your group to be able to join. You can\r\ntell them the name of your group by any means you want, or use the system\'s\r\n"invite" function.\r\n\r\nAgain, the name is:\r\n"{{ room_group_name }}"\r\n\r\nMembers who are still on the waiting list or cancelled cannot join a group.\r\n\r\n* HOW TO MANAGE YOUR GROUP\r\n\r\nYou can always log into the system and view your group\'s status using your\r\nlogin and password.\r\n\r\nYou can remove any members except yourself, and you can delete the entire\r\ngroup.\r\n\r\nAs the person responsible for a group, you\'ll be receiving email every time\r\nsomething changes about your group, like joining or leaving members, etc.\r\n\r\n* IF YOU NEED HELP\r\n\r\nWrite email to ef@eurofurence.org ... but please do this only if you\'re\r\nreally stuck. We\'re expecting many guests, and we\'re doing this in our\r\nspare time :)\r\n\r\n* IMPORTANT\r\n\r\nIf you wish to hand responsibility for your group to someone else, please\r\nwrite to ef@eurofurence.org.\r\n\r\nYours,\r\n\r\n    the EF registration system.', '2021-11-21 12:12:57', '2021-11-21 12:12:57', NULL);
