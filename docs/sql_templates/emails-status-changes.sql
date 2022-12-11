-- currently in use

INSERT INTO `mail_templates` (`id`, `cid`, `lang`, `subject`, `data`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('076943b0-4a32-11ec-b88e-3431c4db8789', 'change-status-new', 'en-US', 'New', 'Dear {{ nickname }},\r\n\r\nthis is an automated verification message from the Eurofurence \r\nregistration system. We have received your registration for Eurofurence, \r\nEurope\'s largest annual furry convention.\r\n\r\nYour status is: NEW - We have received your application\r\n\r\nA space at Eurofurence has NOT yet been reserved for you. You first have \r\nto verify your email address and confirm your registration by clicking \r\non the following link:\r\n\r\n{{ link }}\r\n\r\nOnce you have done this, your registration will be processed further. \r\nYour registration will then be reviewed by our staff, and you should \r\nreceive another mail from us within a couple of days.\r\n\r\nIf you have any questions, feel free to email us. Simply reply to this \r\nmessage. More information on Eurofurence is available at\r\n\r\n       https://www.eurofurence.org/\r\n\r\nYours,\r\n\r\nThe Eurofurence Team', '2021-11-20 19:45:27', '2021-11-21 11:59:10', NULL),
	('0769a5ef-4a32-11ec-b88e-3431c4db8789', 'change-status-new', 'de-DE', 'Neu', 'Hallo {{ nickname }},\r\n\r\ndies ist eine automatische Nachricht des Eurofurence-Buchungssystems. \r\nWir haben deine Anmeldung für Eurofurence erhalten. Eurofurence ist die \r\ngrösste Furry-Convention in Europa.\r\n\r\nDein Status: NEU - Wir haben deine Anmeldung erhalten\r\n\r\nEs wurde noch kein Platz für dich reserviert. Zuerst bestätige bitte \r\ndeine E-mail-Adresse und deine Anmeldung, indem du auf folgenden Link \r\nklickst:\r\n\r\n{{ link }}\r\n\r\nNachdem du deine Anmeldung bestätigt hast, wird dein Antrag weiter \r\nbearbeitet. Einer unserer Mitarbeiter wird sich so schnell wie möglich \r\num deine Anmeldung kümmern, und du solltest innerhab weniger Tage eine \r\nweitere E-Mail erhalten.\r\n\r\nWenn du Fragen hast, schreibe uns eine E-Mail. Es genügt, auf diese Mail \r\nzu antworten. Mehr Informationen über Eurofurence gibt es hier:\r\n\r\n       https://www.eurofurence.org/\r\n\r\nGrüße\r\n\r\nDas Eurofurence Team', '2021-11-20 19:45:27', '2021-11-21 11:59:19', NULL),
    ('38c830d5-4a2c-11ec-b88e-3431c4db8789', 'change-status-approved', 'en-US', 'Accept', 'Dear {{ nickname }},\r\n\r\nThis message is to confirm your registration or an update of your\r\nmembership information. Do not be alarmed if you receive this message\r\nmultiple times, it is sent out to you again every time your registration\r\ninfo is updated.\r\n\r\n====> Status\r\n\r\nYour status is            : PENDING - We are awaiting your payment\r\nYour payment is due until : {{ .DUE_DATE }}\r\n\r\nThis means there is now a place reserved for you at Eurofurence. Please\r\nmake your payments before the due dates mentioned above, otherwise your\r\nregistration will be cancelled.\r\n\r\nPlease make sure that any fees that your particular way of paying may\r\ncause will not be debited to us. If it happens anyway, we will ask you\r\nto pay the difference when you arrive at the convention.\r\n\r\nExact amounts due at each date can be viewed if you log in to the\r\nregistration system and choose "How to pay" from the navigation bar.\r\n\r\n====> Your payment options:\r\n\r\nPlease log in to the registration system and choose "How to pay" from\r\nthe navigation bar to obtain detailed information on payment options,\r\nincluding account data and functionality that will let you perform\r\ncredit card payments.\r\n\r\n====> Booking a hotel room:\r\n\r\nPlease log in to the registration system and choose "How to book a room"\r\nfrom the green navigation bar to the left for instructions.\r\n\r\n====> Changing your personal information:\r\n\r\nIf you wish to make any final changes to your registration, you still\r\nhave the chance. You will find a summary of your account data at the end\r\nof this document, below the German translation. If you want to add\r\npackages or change any of your contact details, please visit\r\n\r\n        https://reg.eurofurence.org/regsys/login.jsp\r\n\r\nand log in to your registration, then select "Edit Registration". You\r\nwill see your registration form, and can fix any settings that you might\r\nhave left out in the initial hurry of getting registered.\r\n\r\nYours,\r\n\r\nThe Eurofurence Team', '2021-11-20 19:03:53', '2021-11-21 11:58:02', NULL),
    ('80316b45-4a2c-11ec-b88e-3431c4db8789', 'change-status-approved', 'de-DE', 'Akzeptieren', 'Hallo {{ nickname }},\r\n\r\nDiese Nachricht ist die Bestätigung deiner Anmeldung, oder einer\r\nÄnderung deiner Mitgliedsdaten. Es kann sein, daß du diese Nachricht\r\nmehrmals bekommst, sie wird jedes mal verschickt, wenn an deinen\r\nAnmeldedaten etwas verändert wurde.\r\n\r\n====> Status\r\n\r\nDein Status     : SCHWEBEND - Wir warten auf deine Zahlung!\r\nZahlungsfrist   : {{ .DUE_DATE }}\r\n\r\nDas bedeutet, es ist auf EF ein Platz für dich reserviert. Bitte zahle\r\ninnerhalb der oben angegebenen Frist. Läuft diese Frist ab ohne daß wir\r\nvon dir gehört haben, wird deine Anmeldung von uns storniert.\r\n\r\nBitte stelle sicher, dass uns der komplette Betrag ohne Abzug von\r\nGebühren auf deiner Seite erreicht. Wenn bei uns zu wenig Geld ankommt,\r\nwerden wir dich bei der Anmeldung vor Ort darum bitten den\r\nDifferenzbetrag nachzuzahlen.\r\n\r\n===> Zahlungsmöglichkeiten:\r\n\r\nLogge dich auf der Registrierseite ein und wähle auf der linken Seite\r\n"How to pay" aus, um genauere Informationen zu erhalten.\r\n\r\n====> Hotelzimmer buchen:\r\n\r\nLogge dich auf der Registrierseite ein und wähle auf der linken Seite\r\n"How to book a room" aus, um genauere Informationen zu erhalten,\r\nwie du ein Zimmer zu den Sonderkonditionen für Eurofurence buchen\r\nkannst.\r\n\r\n====> Anmeldedaten ändern\r\n\r\nEs ist noch möglich, letzte Änderungen an deiner Anmeldung\r\ndurchzuführen. Du findest eine Zusammenfassung deiner Login-Daten ganz\r\nunten am Ende dieser Mail. Wenn du noch Pakete hinzufügen, oder deine\r\nKontaktinformationen ändern möchtest, besuche bitte\r\n\r\n        https://reg.eurofurence.org/regsys/login.jsp\r\n\r\nund logge dich mit deinen untenstehenen Daten ein, und wähle "Edit\r\nRegistration". Du siehst dann eine Übersicht über alle deine\r\nAnmeldedaten und kannst nachträglich alles beheben, was eventuell in der\r\nersten Anmeldehektik schief gegangen sein könnte.\r\n\r\nGrüße\r\n\r\nDas Eurofurence Team', '2021-11-20 19:05:53', '2021-11-21 11:58:05', NULL),
    ('fd822296-4a2d-11ec-b88e-3431c4db8789', 'change-status-partially paid', 'en-US', 'Partial Payment', 'Dear {{ nickname }},\r\n\r\nThis message is to confirm your registration or an update of your \r\nmembership information. Do not be alarmed if you receive this message \r\nmultiple times, it is sent out to you again every time your registration \r\ninfo is updated.\r\n\r\n====> Status\r\n\r\nYour status is            : PARTIALLY PAID\r\nYour payment is due until : {{ .DUE_DATE }}\r\n\r\nPlease make sure that any fees that your particular way of paying may \r\ncause will not be debited to us. If it happens anyway, we will ask you \r\nto pay the difference when you arrive at the con.\r\n\r\n====> Your payment options:\r\n\r\nPlease log in to the registration system and choose "How to pay" from \r\nthe navigation bar to obtain detailed information on payment options, \r\nincluding account data and functionality that will let you perform \r\ncredit card payments.\r\n\r\nYours,\r\n\r\nThe Eurofurence Team', '2021-11-20 19:16:32', '2021-11-21 13:44:42', NULL),
    ('0fee7a7d-4a2e-11ec-b88e-3431c4db8789', 'change-status-partially paid', 'de-DE', 'Teilzahlung', 'Hallo {{ nickname }},\r\n\r\nDiese Nachricht ist die Bestätigung deiner Anmeldung, oder einer \r\nÄnderung deiner Mitgliedsdaten. Es kann sein, daß du diese Nachricht \r\nmehrmals bekommst, sie wird jedes mal verschickt, wenn an deinen \r\nAnmelde-Daten etwas verändert wurde.\r\n\r\n====> Status\r\n\r\nDein Status   : TEILZAHLUNG ERHALTEN\r\nZahlungsfrist : {{ .DUE_DATE }}\r\n\r\nBitte stelle sicher, dass uns der komplette Betrag ohne Abzug von \r\nGebühren auf deiner Seite erreicht. Wenn bei uns zu wenig Geld ankommt, \r\nwerden wir dich bei der Anmeldung vor Ort darum bitten den \r\nDifferenzbetrag nachzuzahlen.\r\n\r\n===> Zahlungsmöglichkeiten:\r\n\r\nLogge dich auf der Registrierseite ein und wähle auf der linken Seite \r\n"How to pay" aus, um genauere Informationen zu erhalten.\r\n\r\nGrüße\r\n\r\nDas Eurofurence Team', '2021-11-20 19:17:03', '2021-11-21 11:58:20', NULL),
    ('1ed6b8cd-4a30-11ec-b88e-3431c4db8789', 'change-status-paid', 'en-US', 'Pay', 'Dear {{ nickname }},\r\n\r\nThis message is to confirm your payment or an update of your membership \r\ninformation. Do not be alarmed if you receive this message multiple \r\ntimes, it is sent out to you again every time your registration info is \r\nupdated.\r\n\r\nYour status is: PAID - Your registration is now officially complete.\r\n\r\nIf you have any questions, feel free to email us. Simply reply to this \r\nmessage. You can edit your registration info by pointing your web \r\nbrowser to\r\n\r\n       https://reg.eurofurence.org/\r\n\r\nand clicking on \'User login\'. \r\n\r\nYou can find a summary of your account data at the end of this document, \r\nbelow the German translation.\r\n\r\nYours,\r\n\r\nThe Eurofurence Team', '2021-11-20 19:31:47', '2021-11-21 11:58:29', NULL),
    ('e54a0bd6-4a30-11ec-b88e-3431c4db8789', 'change-status-paid', 'de-DE', 'Bezahlung', 'Hallo {{ nickname }},\r\n\r\nDiese Nachricht ist die Bestätigung deiner Bezahlung, oder einer \r\nÄnderung deiner Mitgliedsdaten. Es kann sein, daß du diese Nachricht \r\nmehrmals bekommst, sie wird jedes mal verschickt, wenn an deinen \r\nAnmelde-Daten etwas verändert wurde.\r\n\r\nDein Status: BEZAHLT - Deine Anmeldung ist nun offiziell gültig\r\n\r\nWenn es noch Fragen gibt, so beantworten wir diese gerne per E-Mail. Es \r\ngenügt, auf diese Mail zu antworten. Du kannst unter der folgenden URL \r\ndeine persönlichen Daten jederzeit einsehen und verändern:\r\n\r\n       https://reg.eurofurence.org/\r\n\r\nDort einfach auf "User Login" klicken. \r\n\r\nEine Zusammenfassung deiner Account-Daten findest du am Ende dieser \r\nE-Mail, unterhalb der Signatur.\r\n\r\nGrüße\r\n\r\nDas Eurofurence Team', '2021-11-20 19:37:20', '2021-11-21 11:58:32', NULL),
    ('10054ac0-4a2f-11ec-b88e-3431c4db8789', 'change-status-cancelled', 'en-US', 'Cancel', 'Dear {{ nickname }},\r\n\r\nthis is to inform you that your Eurofurence registration has been \r\ncancelled.\r\n\r\nYour status is: CANCELLED - {{ reason }}\r\n\r\nIf you wish to re-apply or if you think that this cancellation is an \r\nerror on our side, please contact us by simply replying to this email.\r\n\r\nYours,\r\n\r\nThe Eurofurence Team', '2021-11-20 19:24:13', '2021-11-21 11:59:00', NULL),
    ('2251007f-4a2f-11ec-b88e-3431c4db8789', 'change-status-cancelled', 'de-DE', 'Abbruch', 'Hallo {{ nickname }},\r\n\r\nleider müssen wir dir mitteilen, daß deine Anmeldung storniert wurde.\r\n\r\nDein Status: STORNIERT - {{ reason }}\r\n\r\nWenn du dich erneut anmelden möchtest, oder wenn du meinst, daß die \r\nStornierung ein Irrtum auf unserer Seite war, dann melde dich bitte bei \r\nuns. Es genügt, einfach auf diese E-Mail zu antworten.\r\n\r\nGrüße\r\n\r\nDas Eurofurence-Team', '2021-11-20 19:24:44', '2021-11-21 11:58:56', NULL),
	('0769ff1c-4a32-11ec-b88e-3431c4db8789', 'change-status-waiting', 'en-US', 'Wait', 'Dear {{ nickname }},\r\n\r\nThis is an automated message to inform you that you have been put on the \r\nEurofurence waiting list.\r\n\r\nYour status: WAITING - Your registration is on hold\r\n\r\nYou\'ve got these options:\r\n\r\n1) Wait until space becomes available\r\n\r\n  Usually at least some already registered members will cancel their\r\n  membership, allowing others to take their place. \r\n\r\n2) Cancel your registration\r\n\r\n  If you decide you cannot wait until a space becomes available, please \r\n  let us know so we can remove you from the waiting list. Thank you!\r\n\r\nIf you have any questions, feel free to email us. All you have to do is \r\nreply to this message. You can edit your registration info by pointing \r\nyour web browser to\r\n\r\n       https://reg.eurofurence.org/\r\n\r\nand clicking on \'User login\'.\r\n\r\nYou can find a summary of your account data at the end of this document, \r\nbelow the German translation.\r\n\r\nYours,\r\n\r\nThe Eurofurence Team', '2021-11-20 19:45:27', '2021-11-21 11:59:08', NULL),
	('076a6299-4a32-11ec-b88e-3431c4db8789', 'change-status-waiting', 'de-DE', 'Warten', 'Hallo {{ nickname }},\r\n\r\ndies ist eine automatische Nachricht um dir mitzuteilen, daß du auf die \r\nEurofurence Warteliste gesetzt wurdest.\r\n\r\nDein Status: WARTELISTE - Deine Anmeldung liegt auf Eis\r\n\r\nEs gibt jetzt folgende Möglichkeiten:\r\n\r\n1) Warten bis ein Platz frei wird\r\n\r\n  Normalerweise gibt es bis zur Convention immer noch eine Menge\r\n  Stornierungen, so daß Leute von der Warteliste nachrücken können. \r\n\r\n2) Abmeldung\r\n\r\n  Wenn du nicht warten kannst, bis ein Platz freiwird, dann melde dich\r\n  bitte bei uns, damit wir deine Anmeldung löschen und dich von der\r\n  Warteliste nehmen können. Danke!\r\n\r\nWenn es noch Fragen gibt, so beantworten wir diese gerne per E-Mail. Es \r\ngenügt, auf diese Mail zu antworten. Du kannst unter der folgenden URL \r\ndeine persönlichen Daten jederzeit einsehen und verändern:\r\n\r\n       https://reg.eurofurence.org/\r\n\r\nDort einfach auf "User Login" klicken.\r\n\r\nEine Zusammenfassung deiner Account-Daten findest du am ende dieser \r\nE-Mail, unterhalb der Signatur.\r\n\r\nGrüße\r\n\r\nDas Eurofurence Team', '2021-11-20 19:45:27', '2021-11-21 11:59:06', NULL),
	('076ac048-4a32-11ec-b88e-3431c4db8789', 'changed_email', 'en-US', 'Changed Email', 'Dear {{ nickname }},\r\n\r\nthis is an automated verification message from the Eurofurence \r\nregistration system. You have requested to change your email address and \r\nwe\'d like to ask you to confirm that this email has reached you by \r\nclicking on the link below:\r\n\r\nNew address: {{ new_email }}\r\n{{ link }}\r\n\r\nUntil you verified this email address, all Eurofurence mail will \r\ncontinue to be sent to the previously verified address.\r\n\r\nOld address: {{ email }}\r\n\r\nIf you have any questions, feel free to email us. Simply reply to this \r\nmessage. More information on Eurofurence is available at\r\n\r\n       https://www.eurofurence.org/\r\n\r\nYours,\r\n\r\nThe Eurofurence Team', '2021-11-20 19:45:27', '2021-11-21 11:59:03', NULL),
	('2e9d5528-4a32-11ec-b88e-3431c4db8789', 'changed_email', 'de-DE', 'Email Geändert', 'Hallo {{ nickname }},\r\n\r\ndies ist eine automatische Nachricht des Eurofurence-Buchungssystems. Du \r\nhast deine Email-Adresse geändert und wir möchten dich bitten, den \r\nEmpfang dieser Email zu bestätigen indem du auf den folgenden Link \r\nklickst:\r\n\r\nNeue Adresse: {{ new_email }}\r\n{{ link }}\r\n\r\nBis du diese neue Adresse bestätigt hast werden alle Eurofurence Mails \r\nweiterhin an deine vorherige, bestätigte Adresse geschickt.\r\n\r\nAlte Adresse: {{ email }} \r\n\r\nWenn du Fragen hast, schreibe uns eine E-Mail. Es genügt, auf diese Mail \r\nzu antworten. Mehr Informationen über Eurofurence gibt es hier:\r\n\r\n       https://www.eurofurence.org/\r\n\r\nGrüße\r\n\r\nDas Eurofurence Team', '2021-11-20 19:46:33', '2021-11-21 11:59:04', NULL),
	('85255fc0-4a2f-11ec-b88e-3431c4db8789', 'info', 'en-US', 'Information', 'Registration info:\r\n------------------\r\n\r\nMembership Number                :   {{ badge_number }}\r\nNickname                         :   {{ nickname }}\r\nTotal amount due                 :   {{ total_dues }}\r\nDues Remaining                   :   {{ remaining_dues }}\r\n\r\nTo log into the registration system, go to\r\n\r\n        https://reg.eurofurence.org/regsys/login.jsp\r\n \r\nIf you have forgotten your password, you can use the \'Forgot password?\'\r\nlink on this page to reset your password.', '2021-11-20 19:27:29', '2021-11-21 13:45:31', NULL),
	('9f1b7827-4a2f-11ec-b88e-3431c4db8789', 'info', 'de-DE', 'Information', 'Registrations Information:\r\n------------------\r\n\r\nRegistriernummer                 :   {{ badge_number }}\r\nNickname                         :   {{ nickname }}\r\nGesamtpreis                      :   {{ total_dues }}\r\nOffener Betrag                   :   {{ remaining_dues }}\r\n\r\nTo log into the registration system, go to\r\n\r\n        https://reg.eurofurence.org/regsys/login.jsp\r\n \r\nIf you have forgotten your password, you can use the \'Forgot password?\'\r\nlink on this page to reset your password.', '2021-11-20 19:28:13', '2021-11-21 11:58:09', NULL),
    ('85e1822b-4a2d-11ec-b88e-3431c4db8789', 'guest', 'en-US', 'Guest', 'Hello and welcome to Eurofurence!\r\n\r\nIt is our pleasure to inform you have been registered as a special guest \r\nof the convention. Among other things, this means no con fee, free \r\nhousing at the con, access to all areas, and supersponsor privileges.\r\n\r\nYour status is: GUEST - No further actions required\r\n\r\nIf you have any questions, feel free to email us. Simply reply to this \r\nmessage. You can edit your registration info by pointing your web \r\nbrowser to\r\n\r\n       https://reg.eurofurence.org/\r\n\r\nand clicking on \'User login\'.\r\n\r\nYou can find a summary of your account data at the end of this document \r\nbelow the German translation.\r\n\r\nYours,\r\n\r\nThe Eurofurence Team', '2021-11-20 19:13:12', '2021-11-21 11:57:26', NULL),
    ('9b7585fd-4a2d-11ec-b88e-3431c4db8789', 'guest', 'de-DE', 'Gast', 'Hallo und willkommen zu Eurofurence!\r\n\r\nEs ist uns eine Freude Ihnen mitzuteilen, daß Sie als "Special Guest" \r\nder Convention angemeldet wurden. Das bedeutet: Kein Mitgliedsbeitrag, \r\nkostenlose Unterbringung, Zugang zu allen Bereichen und \r\nSupersponsor-Privilegien.\r\n\r\nStatus: GAST - Keine weiteren Aktionen nötig\r\n\r\nWenn es noch Fragen gibt, so beantworten wir diese gerne per E-Mail. Es \r\ngenügt, auf diese Mail zu antworten. Sie können unter der folgenden URL \r\nIhre persönlichen Daten jederzeit einsehen und verändern:\r\n\r\n       https://reg.eurofurence.org/\r\n\r\nDort einfach auf "User Login" klicken.\r\n\r\nEine Zusammenfassung Ihrer Account-Daten finden Sie am Ende dieser \r\nE-Mail, unterhalb der Signatur.\r\n\r\nGrüße\r\n\r\nDas Eurofurence Team', '2021-11-20 19:13:48', '2021-11-21 11:57:29', NULL),
	('e54a67c9-4a30-11ec-b88e-3431c4db8789', 'remind', 'en-US', 'Remind', 'Dear {{ nickname }},\r\n\r\nAccording to our database you have not yet paid one or more Eurofurence \r\nmembership packages. This payment was due by {{ .DUE_DATE }}.\r\n\r\nIf we do not receive payment or hear from you otherwise within the next \r\nseven days, your registration will be cancelled.\r\n\r\n====> Status\r\n\r\nYour status is: OVERDUE - Please pay your outstanding fees.\r\n\r\nIf you find you cannot attend, please reply to this mail and tell us, so \r\nyour space can be reassigned. If you find that you\'ve received this \r\nnotice on behalf of an error on our side, please also contact us \r\nimmediately.\r\n\r\n====> Your payment options:\r\n\r\nPlease log in to the registration system and choose "How to pay" from \r\nthe navigation bar to obtain detailed information on payment options, \r\nincluding account data and functionality that will let you perform \r\ncredit card payments.\r\n\r\n====> Any questions?\r\n\r\nIf you have any questions, feel free to email us. Simply reply to this \r\nmessage. You can edit your registration info by pointing your web \r\nbrowser to\r\n\r\n       https://reg.eurofurence.org/\r\n\r\nand clicking on \'User login\'.\r\n\r\nYou can find a summary of your account data at the end of this document, \r\nbelow the German translation.\r\n\r\nYours,\r\n\r\nThe Eurofurence Team', '2021-11-20 19:37:20', '2021-11-21 11:57:18', NULL),
	('e54ad4ab-4a30-11ec-b88e-3431c4db8789', 'remind', 'de-DE', 'Erinnerung', 'Hallo {{ nickname }},\r\n\r\nlaut unserer Datenbank haben wir von dir noch nicht die volle Bezahlung \r\ndeiner Eurofurence-Teilnehmergebühren erhalten. Die Zahlung war fällig \r\nzum {{ .DUE_DATE }}.\r\n\r\nWenn wir in den naechsten 7 Tagen von dir weder eine Rückmeldung noch \r\neine Zahlung erhalten, werden wir deine Anmeldung stornieren.\r\n\r\n====> Status\r\n\r\nDein Status: ÜBERFÄLLIG - Bitte Beitrag bezahlen, sonst Storno!\r\n\r\nSollte sich inzwischen herausgestellt haben daß du nicht teilnehmen \r\nkannst, melde dich bitte bei uns, so daß wir deinen Platz anderweitig \r\nvergeben können. Sollte es sich bei dieser Mahnung um einen Irrtum \r\nhandeln, dann melde dich bitte ebenfalls so schnell wie möglich.\r\n\r\n===> Zahlungsmöglichkeiten:\r\n\r\nLogge dich auf der Registrierseite ein und wähle auf der linken Seite \r\n"How to pay" aus, um genauere Informationen zu erhalten.\r\n\r\n====> Irgendwelche Fragen?\r\n\r\nWenn es noch Fragen gibt, so beantworten wir diese gerne per E-Mail. Es \r\ngenügt, auf diese Mail zu antworten. Du kannst unter der folgenden URL \r\ndeine persönlichen Daten jederzeit einsehen und verändern:\r\n\r\n       https://reg.eurofurence.org/\r\n\r\nDort einfach auf "User Login" klicken.\r\n\r\nEine Zusammenfassung deiner Account-Daten findest du am Ende dieser \r\nE-Mail, unterhalb der Signatur.\r\n\r\nGrüße\r\n\r\nDas Eurofurence Team', '2021-11-20 19:37:20', '2021-11-21 11:57:20', NULL);

