package sys

import (
	"reflect"
	"testing"
)

func TestSubtract(t *testing.T) {
	CalibratedStatePaths := []string{
		"c:\\users\\mohd zeeshan\\appdata\\local\\programs\\microsoft vs code\\code.exe",
		"c:\\program files (x86)\\microsoft\\edge\\application\\msedge.exe",
		"c:\\windows\\systemapps\\microsoftwindows.client.cbs_cw5n1h2txyewy\\textinputhost.exe",
		"c:\\program files\\windowsapps\\microsoft.windowsterminal_1.23.20211.0_x64__8wekyb3d8bbwe\\windowsterminal.exe",
	}

	aStatePaths := []string{
		"c:\\users\\mohd zeeshan\\appdata\\local\\programs\\microsoft vs code\\code.exe",
		"c:\\program files (x86)\\microsoft\\edge\\application\\msedge.exe",
		"c:\\windows\\systemapps\\microsoftwindows.client.cbs_cw5n1h2txyewy\\textinputhost.exe",
		"c:\\program files\\windowsapps\\microsoft.windowsterminal_1.23.20211.0_x64__8wekyb3d8bbwe\\windowsterminal.exe",
		"c:\\program files (x86)\\videolan\\vlc\\vlc.exe",
		"c:\\program files\\windowsapps\\5319275a.whatsappdesktop_2.2587.10.0_x64__cv1g1gvanyjgm\\whatsapp.root.exe",
	}

	got := Subtract(aStatePaths, CalibratedStatePaths)

	want := []string{
		"c:\\program files (x86)\\videolan\\vlc\\vlc.exe",
		"c:\\program files\\windowsapps\\5319275a.whatsappdesktop_2.2587.10.0_x64__cv1g1gvanyjgm\\whatsapp.root.exe",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n %v \n want:\n %v\n", got, want)

	}
}
