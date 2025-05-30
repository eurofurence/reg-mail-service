package templates

import (
	"context"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/database/dbrepo"
	"github.com/google/uuid"
)

// SeedDefaultTemplates provides the minimum required default templates to seed the database.
//
// This is particularly useful for the in-memory database, which otherwise would lack templates needed
// for system operation.
//
// If your database already contains the required templates, they are not touched.
func SeedDefaultTemplates(ctx context.Context, db dbrepo.Repository) error {
	existing, err := db.GetTemplates(ctx)
	if err != nil {
		return err
	}

	for _, lang := range defaultLanguages {
		for _, tpl := range defaultTemplates {
			alreadyExists := false
			for _, ex := range existing {
				if tpl.CommonID == ex.CommonID && tpl.Language == ex.Language {
					alreadyExists = true
				}
			}

			if !alreadyExists {
				newId, _ := uuid.NewUUID()
				tpl.ID = newId.String()
				tpl.Language = lang

				aulogging.Logger.NoCtx().Info().Printf("creating required template %s %s with subject '%s' and default text", tpl.CommonID, tpl.Language, tpl.Subject)
				err := db.CreateTemplate(ctx, &tpl)
				if err != nil {
					aulogging.Logger.NoCtx().Error().WithErr(err).Printf("failed to create required template %s %s: %s", tpl.CommonID, tpl.Language, err.Error())
				}
			}
		}
	}

	return nil
}

var defaultLanguages = []string{"en-US", "de-DE"}

var defaultTemplates []entity.Template = []entity.Template{
	{
		CommonID: "change-status-approved",
		Subject:  "Registration Confirmed - Please Pay",
		Data: `Dear {{ nickname }},

This message is to confirm your registration or an update of your attendance information.

====> Status

Your status is            : PENDING - We are awaiting your payment
Your payment is due until : {{ due_date }}

Please make your payments before the due date mentioned above, otherwise your registration will be cancelled.

Yours,

The Registration Team


Registration info:
------------------

Registration ID                  :   {{ badge_number }}
Nickname                         :   {{ nickname }}
Total amount due                 :   {{ total_dues }}
Dues Remaining                   :   {{ remaining_dues }}
`,
	},
	{
		CommonID: "change-status-cancelled",
		Subject:  "Registration Cancelled",
		Data: `Dear {{ nickname }},

this is to inform you that your registration has been cancelled.

Your status is: CANCELLED - {{ reason }}

If you wish to re-apply or if you think that this cancellation is an error on our side, please contact us by simply replying to this email.

Yours,

The Registration Team`,
	},
	{
		CommonID: "change-status-new",
		Subject:  "New Registration",
		Data: `Dear {{ nickname }},

this is an automated verification message from the registration system.

Your status is: NEW - We have received your application

Your registration will be reviewed by our staff, and you should receive another mail from us when things are ready.

If you have any questions, feel free to email us. Simply reply to this message.

Yours,

The Registration Team
`,
	},
	{
		CommonID: "change-status-paid",
		Subject:  "Registration Paid",
		Data: `Dear {{ nickname }},

This message is to confirm your payment.

Your status is: PAID - Your registration is now officially complete.

If you have any questions, feel free to email us. Simply reply to this message. You can edit your registration info by pointing your web browser to

       {{ regsys_url }}

You can find a summary of your information at the end of this document.

Yours,

The Registration Team


Registration info:
------------------

Redistration ID                  :   {{ badge_number }}
Nickname                         :   {{ nickname }}
Total amount due                 :   {{ total_dues }}
Dues Remaining                   :   {{ remaining_dues }}
`,
	},
	{
		CommonID: "change-status-partially paid",
		Subject:  "Partial Payment - Please Pay Remaining Amount",
		Data: `Dear {{ nickname }},

This message is to confirm your registration or an update of your attendance information.

====> Status

Your status is            : PARTIALLY PAID
Your payment is due until : {{ due_date }}

Please make sure that any fees that your particular way of paying may cause will not be debited to us.

Yours,

The Registration Team


Registration info:
------------------

Registration ID                  :   {{ badge_number }}
Nickname                         :   {{ nickname }}
Total amount due                 :   {{ total_dues }}
Dues Remaining                   :   {{ remaining_dues }}
`,
	},
	{
		CommonID: "change-status-waiting",
		Subject:  "On the Waiting List",
		Data: `Dear {{ nickname }},

This is an automated message to inform you that you have been put on the waiting list.

Your status: WAITING - Your registration is on hold

You've got these options:

1) Wait until space becomes available

   Usually at least some people will cancel, allowing others to take their place.

2) Cancel your registration

   If you decide you cannot wait until a space becomes available, please let us know so we can remove you from the waiting list. Thank you!

If you have any questions, feel free to email us. All you have to do is reply to this message. 

You can edit your registration info by pointing your web browser to

       {{ regsys_url }}

You can find a summary of your information at the end of this document.

Yours,

The Registration Team


Registration info:
------------------

Registration ID                  :   {{ badge_number }}
Nickname                         :   {{ nickname }}
`,
	},
	{
		CommonID: "guest",
		Subject:  "Guest of the Convention",
		Data: `Hello and welcome!

It is our pleasure to inform you have been registered as a special guest of the convention. Among other things, this means no con fee, free housing at the con, access to all areas, and supersponsor privileges.

Your status is: GUEST - No further actions required

If you have any questions, feel free to email us. Simply reply to this message.

You can edit your registration info by pointing your web browser to

       {{ regsys_url }}

You can find a summary of your information at the end of this document.

Yours,

The Registration Team


Registration info:
------------------

Registration ID                  :   {{ badge_number }}
Nickname                         :   {{ nickname }}
`,
	},
	{
		CommonID: "payment-cncrd-adapter-error",
		Subject:  "Payment Adapter Error Notice",
		Data: `Encountered an unexpected condition during {{ operation }}

ReferenceId: {{ referenceId }}
Status:      {{ status }}

Please have a look at the logs for details.
`,
	},
	// roomshare groups: mails to group owner
	//
	// nickname, groupname, object_badge_number, object_nickname
	{
		CommonID: "group-member-joined",
		Subject:  "New group member",
		Data: `Dear {{ nickname }}!

A furry named

{{ object_nickname }} ({{ object_badge_number }})

has joined your group

{{ groupname }}

and wants to be one of your roommates.

If you do not wish to have that person in your group, you can log into the registration system and remove them. An automated note will be sent to that person if you do this.

Yours,

The Registration Team
`,
	},
	{
		CommonID: "group-member-applied",
		Subject:  "New group application",
		Data: `Dear {{ nickname }}!

A furry named

{{ object_nickname }} ({{ object_badge_number }})

has applied to your group

{{ groupname }}

and wants to be one of your roommates.

Please log into the registration system and either accept or decline the application.

If you decline, you will also have the option to prevent this furry from applying again.

Yours,

The Registration Team
`,
	},
	{
		CommonID: "group-member-removed",
		Subject:  "Group member was removed",
		Data: `Dear {{ nickname }}!

A furry named

{{ object_nickname }} ({{ object_badge_number }})

has been removed from your group

{{ groupname }}

by an administrator. This may also have happened automatically because their registration has been cancelled.

Yours,

The Registration Team
`,
	},
	{
		CommonID: "group-invitation-declined",
		Subject:  "Your group invitation was declined",
		Data: `Dear {{ nickname }}!

A furry named

{{ object_nickname }} ({{ object_badge_number }})

has declined your invitation to join your group

{{ groupname }}.

Unless you are sure this was done by mistake, please do not invite them again. Warning: Repeatedly inviting someone to a group they do not wish to be in may in extreme cases get you banned from the convention for disruptive behaviour.

Yours,

The Registration Team
`,
	},
	{
		CommonID: "group-member-left",
		Subject:  "Group member left your group",
		Data: `Dear {{ nickname }}!

A furry named

{{ object_nickname }} ({{ object_badge_number }})

has left your group

{{ groupname }}.

Unless you are sure this was done by mistake, please do not invite them again. Warning: Repeatedly inviting someone to a group they do not wish to be in may in extreme cases get you banned from the convention for disruptive behaviour.

Yours,

The Registration Team
`,
	},
	// roomshare groups: mails to group members
	//
	// nickname, groupname, owner, url (join link)
	{
		CommonID: "group-invited",
		Subject:  "You have been invited to join a room sharing group",
		Data: `Dear {{ nickname }}!

A furry called

{{ owner }}

has invited you to join their room sharing group

{{ groupname }}

and wants you to be one of their roommates.

You can accept or decline the invitation by clicking on this link:

{{ url }}

If you decline, you will also have the option to prevent this furry from inviting you again.

Yours,

The Registration Team
`,
	},
	{
		CommonID: "group-application-accepted",
		Subject:  "You have been added to a room sharing group",
		Data: `Dear {{ nickname }}!

You have been added to the room sharing group

{{ groupname }}

owned by a furry called

{{ owner }}

If you change your mind about being in the group, you can log into the registration system and leave it.

Yours,

The Registration Team
`,
	},
	{
		CommonID: "group-member-kicked",
		Subject:  "You have been removed from a room sharing group",
		Data: `Dear {{ nickname }}!

You have been removed from the room sharing group

{{ groupname }}.

This is normally done by the creator of the group, or in very rare cases by the admins of the registration system.

Please do not simply add yourself again, but rather clear up the situation with the group creator. Warning: Repeatedly adding yourself to groups whose creators do not wish to have you in there may in extreme cases get you banned from the convention for disruptive behaviour.

Yours,

The Registration Team
`,
	},
	{
		CommonID: "group-application-declined",
		Subject:  "Room sharing group application declined",
		Data: `Dear {{ nickname }}!

Your application for the room sharing group

{{ groupname }}

has been declined.

This is normally done by the creator of the group, or in very rare cases by the admins of the registration system.

Please do not simply add yourself again, but rather clear up the situation with the group creator. Warning: Repeatedly adding yourself to groups whose creators do not wish to have you in there may in extreme cases get you banned from the convention for disruptive behaviour.

Yours,

The Registration Team
`,
	},
	{
		CommonID: "group-new-owner",
		Subject:  "You are now the owner of a room sharing group",
		Data: `Dear {{ nickname }}!

A furry called

{{ owner }}

has made you the new owner of their room sharing group

{{ groupname }}.

Please log into the registration system if you wish to manage your group. 

As group owner, you will be able to invite others, accept and decline pending applications, and remove group members, as well as rename the group.

Yours,

The Registration Team
`,
	},
}
