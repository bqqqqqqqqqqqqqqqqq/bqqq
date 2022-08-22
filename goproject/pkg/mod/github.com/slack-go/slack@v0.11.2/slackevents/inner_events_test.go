package slackevents

import (
	"encoding/json"
	"testing"
)

func TestAppMention(t *testing.T) {
	rawE := []byte(`
			{
				"type": "app_mention",
				"user": "U061F7AUR",
				"text": "<@U0LAN0Z89> is it everything a river should be?",
				"ts": "1515449522.000016",
				"thread_ts": "1515449522.000016",
				"channel": "C0LAN2Q65",
				"event_ts": "1515449522000016",
				"source_team": "T3MQV36V7",
				"user_team": "T3MQV36V7",
				"blah": "test"
		}
	`)
	err := json.Unmarshal(rawE, &AppMentionEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestAppUninstalled(t *testing.T) {
	rawE := []byte(`
		{
			"type": "app_uninstalled"
		}
	`)
	err := json.Unmarshal(rawE, &AppUninstalledEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestGridMigrationFinishedEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "grid_migration_finished",
				"enterprise_id": "EXXXXXXXX"
			}
	`)
	err := json.Unmarshal(rawE, &GridMigrationFinishedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestGridMigrationStartedEvent(t *testing.T) {
	rawE := []byte(`
			{
				"token": "XXYYZZ",
				"team_id": "TXXXXXXXX",
				"api_app_id": "AXXXXXXXXX",
				"event": {
						"type": "grid_migration_started",
						"enterprise_id": "EXXXXXXXX"
				},
				"type": "event_callback",
				"event_id": "EvXXXXXXXX",
				"event_time": 1234567890
		}
	`)
	err := json.Unmarshal(rawE, &GridMigrationStartedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestLinkSharedEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "link_shared",
				"channel": "Cxxxxxx",
				"user": "Uxxxxxxx",
				"message_ts": "123456789.9875",
				"thread_ts": "123456789.9876",
				"links":
						[
								{
										"domain": "example.com",
										"url": "https://example.com/12345"
								},
								{
										"domain": "example.com",
										"url": "https://example.com/67890"
								},
								{
										"domain": "another-example.com",
										"url": "https://yet.another-example.com/v/abcde"
								}
						]
		}
	`)
	err := json.Unmarshal(rawE, &LinkSharedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestLinkSharedComposerEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "link_shared",
				"channel": "COMPOSER",
				"is_bot_user_member": true,
				"user": "Uxxxxxxx",
				"message_ts": "Uxxxxxxx-909b5454-75f8-4ac4-b325-1b40e230bbd8-gryl3kb80b3wm49ihzoo35fyqoq08n2y",
				"unfurl_id": "Uxxxxxxx-909b5454-75f8-4ac4-b325-1b40e230bbd8-gryl3kb80b3wm49ihzoo35fyqoq08n2y",
				"source": "composer",
				"links": [
					{
						"domain": "example.com",
						"url": "https://example.com/12345"
					},
					{
						"domain": "example.com",
						"url": "https://example.com/67890"
					},
					{
						"domain": "another-example.com",
						"url": "https://yet.another-example.com/v/abcde"
					}
				]
			}
	`)
	err := json.Unmarshal(rawE, &LinkSharedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestMessageEvent(t *testing.T) {
	rawE := []byte(`
			{
				"client_msg_id": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
				"type": "message",
				"channel": "G024BE91L",
				"user": "U2147483697",
				"text": "Live long and prospect.",
				"ts": "1355517523.000005",
				"event_ts": "1355517523.000005",
				"channel_type": "channel",
				"source_team": "T3MQV36V7",
				"user_team": "T3MQV36V7",
				"message": {
					"text": "To infinity and beyond.",
					"edited": {
						"user": "U2147483697",
						"ts": "1355517524.000000"
					}
				},
				"previous_message": {
					"text": "Live long and prospect."
				}
		}
	`)
	err := json.Unmarshal(rawE, &MessageEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestBotMessageEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "message",
				"subtype": "bot_message",
				"ts": "1358877455.000010",
				"text": "Pushing is the answer",
				"bot_id": "BB12033",
				"username": "github",
				"icons": {}
		}
	`)
	err := json.Unmarshal(rawE, &MessageEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestThreadBroadcastEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "message",
				"subtype": "thread_broadcast",
				"channel": "G024BE91L",
				"user": "U2147483697",
				"text": "Live long and prospect.",
				"ts": "1355517523.000005",
				"event_ts": "1355517523.000005",
				"channel_type": "channel",
				"source_team": "T3MQV36V7",
				"user_team": "T3MQV36V7",
				"message": {
					"text": "To infinity and beyond.",
					"root": {
						"text": "To infinity and beyond.",
						"ts": "1355517523.000005"
					},
					"edited": {
						"user": "U2147483697",
						"ts": "1355517524.000000"
					}
				},
				"previous_message": {
					"text": "Live long and prospect."
				}
		}
	`)

	var me MessageEvent
	if err := json.Unmarshal(rawE, &me); err != nil {
		t.Error(err)
	}

	if me.Root != nil {
		t.Error("me.Root should be nil")
	}

	if me.Message.Root == nil {
		t.Fatal("me.Message.Root is nil")
	}

	if me.Message.Root.TimeStamp != "1355517523.000005" {
		t.Errorf("me.Message.Root.TimeStamp = %q, want %q", me.Root.TimeStamp, "1355517523.000005")
	}
}

func TestMemberJoinedChannelEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "member_joined_channel",
				"user": "W06GH7XHN",
				"channel": "C0698JE0H",
				"channel_type": "C",
				"team": "T024BE7LD",
				"inviter": "U123456789"
		}
	`)
	err := json.Unmarshal(rawE, &MemberJoinedChannelEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestMemberLeftChannelEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "member_left_channel",
				"user": "W06GH7XHN",
				"channel": "C0698JE0H",
				"channel_type": "C",
				"team": "T024BE7LD"
		}
	`)
	err := json.Unmarshal(rawE, &MemberLeftChannelEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestPinAdded(t *testing.T) {
	rawE := []byte(`
			{
				"type": "pin_added",
				"user": "U061F7AUR",
				"item": {
					"type": "message",
					"channel":"C0LAN2Q65",
					"message":{
						"type":"message",
						"user":"U061F7AUR",
						"text": "<@U0LAN0Z89> is it everything a river should be?",
						"ts":"1539904112.000100",
						"pinned_to":["C0LAN2Q65"],
						"replace_original":false,
						"delete_original":false
					}
				},
				"channel_id":"C0LAN2Q65",
				"event_ts": "1515449522000016"
		}
	`)
	err := json.Unmarshal(rawE, &PinAddedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestPinRemoved(t *testing.T) {
	rawE := []byte(`
			{
				"type": "pin_removed",
				"user": "U061F7AUR",
				"item": {
					"type": "message",
					"channel":"C0LAN2Q65",
					"message":{
						"type":"message",
						"user":"U061F7AUR",
						"text": "<@U0LAN0Z89> is it everything a river should be?",
						"ts":"1539904112.000100",
						"pinned_to":["C0LAN2Q65"],
						"replace_original":false,
						"delete_original":false
					}
				},
				"channel_id":"C0LAN2Q65",
				"event_ts": "1515449522000016"
		}
	`)
	err := json.Unmarshal(rawE, &PinRemovedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestTokensRevoked(t *testing.T) {
	rawE := []byte(`
	{
		"type": "tokens_revoked",
		"tokens": {
				"oauth": [
						"OUXXXXXXXX"
				],
				"bot": [
						"BUXXXXXXXX"
				]
		}
	}
`)
	tre := TokensRevokedEvent{}
	err := json.Unmarshal(rawE, &tre)
	if err != nil {
		t.Error(err)
	}

	if tre.Type != "tokens_revoked" {
		t.Fail()
	}

	if len(tre.Tokens.Bot) != 1 || tre.Tokens.Bot[0] != "BUXXXXXXXX" {
		t.Fail()
	}

	if len(tre.Tokens.Oauth) != 1 || tre.Tokens.Oauth[0] != "OUXXXXXXXX" {
		t.Fail()
	}
}

func TestEmojiChanged(t *testing.T) {
	var (
		ece EmojiChangedEvent
		err error
	)

	// custom emoji added event
	rawAddE := []byte(`
	{
		"type": "emoji_changed",
		"subtype": "add",
		"name": "picard_facepalm",
		"value": "https://my.slack.com/emoji/picard_facepalm/db8e287430eaa459.gif",
		"event_ts" : "1361482916.000004"
	}
`)
	ece = EmojiChangedEvent{}
	err = json.Unmarshal(rawAddE, &ece)
	if err != nil {
		t.Error(err)
	}
	if ece.Subtype != "add" {
		t.Fail()
	}
	if ece.Name != "picard_facepalm" {
		t.Fail()
	}

	// emoji removed event
	rawRemoveE := []byte(`
	{
		"type": "emoji_changed",
		"subtype": "remove",
		"names": ["picard_facepalm"],
		"event_ts" : "1361482916.000004"
	}
`)
	ece = EmojiChangedEvent{}
	err = json.Unmarshal(rawRemoveE, &ece)
	if err != nil {
		t.Error(err)
	}
	if ece.Subtype != "remove" {
		t.Fail()
	}
	if len(ece.Names) != 1 {
		t.Fail()
	}
	if ece.Names[0] != "picard_facepalm" {
		t.Fail()
	}

	// custom emoji rename event
	rawRenameE := []byte(`
	{
		"type": "emoji_changed",
		"subtype": "rename",
		"old_name": "grin",
		"new_name": "cheese-grin",
		"value": "https://my.slack.com/emoji/picard_facepalm/db8e287430eaa459.gif",
		"event_ts" : "1361482916.000004"
	}
`)
	ece = EmojiChangedEvent{}
	err = json.Unmarshal(rawRenameE, &ece)
	if err != nil {
		t.Error(err)
	}
	if ece.Subtype != "rename" {
		t.Fail()
	}
	if ece.OldName != "grin" {
		t.Fail()
	}
	if ece.NewName != "cheese-grin" {
		t.Fail()
	}
}

func TestWorkflowStepExecute(t *testing.T) {
	// see: https://api.slack.com/events/workflow_step_execute
	rawE := []byte(`
	{
		"type":"workflow_step_execute",
		"callback_id":"open_ticket",
		"workflow_step":{
			"workflow_step_execute_id":"1036669284371.19077474947.c94bcf942e047298d21f89faf24f1326",
			"workflow_id":"123456789012345678",
			"workflow_instance_id":"987654321098765432",
			"step_id":"12a345bc-1a23-4567-8b90-1234a567b8c9",
			"inputs":{
				"example-select-input":{
					"value": "value-two",
					"skip_variable_replacement": false
				}
			},
			"outputs":[
			]
		},
		"event_ts":"1643290847.766536"
	}
	`)

	wse := WorkflowStepExecuteEvent{}
	err := json.Unmarshal(rawE, &wse)
	if err != nil {
		t.Error(err)
	}

	if wse.Type != "workflow_step_execute" {
		t.Fail()
	}
	if wse.CallbackID != "open_ticket" {
		t.Fail()
	}
	if wse.WorkflowStep.WorkflowStepExecuteID != "1036669284371.19077474947.c94bcf942e047298d21f89faf24f1326" {
		t.Fail()
	}
	if wse.WorkflowStep.WorkflowID != "123456789012345678" {
		t.Fail()
	}
	if wse.WorkflowStep.WorkflowInstanceID != "987654321098765432" {
		t.Fail()
	}
	if wse.WorkflowStep.StepID != "12a345bc-1a23-4567-8b90-1234a567b8c9" {
		t.Fail()
	}
	if len(*wse.WorkflowStep.Inputs) == 0 {
		t.Fail()
	}
	if inputElement, ok := (*wse.WorkflowStep.Inputs)["example-select-input"]; ok {
		if inputElement.Value != "value-two" {
			t.Fail()
		}
		if inputElement.SkipVariableReplacement != false {
			t.Fail()
		}
	}
}
