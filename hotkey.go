package hotkey

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

const ModNone = 0
const (
	ModAlt = 1 << iota
	ModCtrl
	ModShift
	ModWin
)

const (
	VK_LWIN                = 91  // Left Windows key (Microsoft® Natural® keyboard)
	VK_RWIN                = 92  // Right Windows key (Natural keyboard)
	VK_APPS                = 93  // Applications key (Natural keyboard)
	VK_SLEEP               = 95  // Computer Sleep key
	VK_NUMPAD0             = 96  // Numeric keypad 0 key
	VK_NUMPAD1             = 97  // Numeric keypad 1 key
	VK_NUMPAD2             = 98  // Numeric keypad 2 key
	VK_NUMPAD3             = 99  // Numeric keypad 3 key
	VK_NUMPAD4             = 100 // Numeric keypad 4 key
	VK_NUMPAD5             = 101 // Numeric keypad 5 key
	VK_NUMPAD6             = 102 // Numeric keypad 6 key
	VK_NUMPAD7             = 103 // Numeric keypad 7 key
	VK_NUMPAD8             = 104 // Numeric keypad 8 key
	VK_NUMPAD9             = 105 // Numeric keypad 9 key
	VK_MULTIPLY            = 106 // Multiply key
	VK_ADD                 = 107 // Add key
	VK_SEPARATOR           = 108 // Separator key
	VK_SUBTRACT            = 109 // Subtract key
	VK_DECIMAL             = 110 // Decimal key
	VK_DIVIDE              = 111 // Divide key
	VK_F1                  = 112 // F1 key
	VK_F2                  = 113 // F2 key
	VK_F3                  = 114 // F3 key
	VK_F4                  = 115 // F4 key
	VK_F5                  = 116 // F5 key
	VK_F6                  = 117 // F6 key
	VK_F7                  = 118 // F7 key
	VK_F8                  = 119 // F8 key
	VK_F9                  = 120 // F9 key
	VK_F10                 = 121 // F10 key
	VK_F11                 = 122 // F11 key
	VK_F12                 = 123 // F12 key
	VK_F13                 = 124 // F13 key
	VK_F14                 = 125 // F14 key
	VK_F15                 = 126 // F15 key
	VK_F16                 = 127 // F16 key
	VK_NUMLOCK             = 144 // NUM LOCK key
	VK_SCROLL              = 145 // SCROLL LOCK key
	VK_LSHIFT              = 160 // Left SHIFT key
	VK_RSHIFT              = 161 // Right SHIFT key
	VK_LCONTROL            = 162 // Left CONTROL key
	VK_RCONTROL            = 163 // Right CONTROL key
	VK_LMENU               = 164 // Left MENU key
	VK_RMENU               = 165 // Right MENU key
	VK_BROWSER_BACK        = 166 // Windows 2000: Browser Back key
	VK_BROWSER_FORWARD     = 167 // Windows 2000: Browser Forward key
	VK_BROWSER_REFRESH     = 168 // Windows 2000: Browser Refresh key
	VK_BROWSER_STOP        = 169 // Windows 2000: Browser Stop key
	VK_BROWSER_SEARCH      = 170 // Windows 2000: Browser Search key
	VK_BROWSER_FAVORITES   = 171 // Windows 2000: Browser Favorites key
	VK_BROWSER_HOME        = 172 // Windows 2000: Browser Start and Home key
	VK_VOLUME_MUTE         = 173 // Windows 2000: Volume Mute key
	VK_VOLUME_DOWN         = 174 // Windows 2000: Volume Down key
	VK_VOLUME_UP           = 175 // Windows 2000: Volume Up key
	VK_MEDIA_NEXT_TRACK    = 176 // Windows 2000: Next Track key
	VK_MEDIA_PREV_TRACK    = 177 // Windows 2000: Previous Track key
	VK_MEDIA_STOP          = 178 // Windows 2000: Stop Media key
	VK_MEDIA_PLAY_PAUSE    = 179 // Windows 2000: Play/Pause Media key
	VK_LAUNCH_MAIL         = 180 // Windows 2000: Start Mail key
	VK_LAUNCH_MEDIA_SELECT = 181 // Windows 2000: Select Media key
	VK_LAUNCH_APP1         = 182 // Windows 2000: Start Application 1 key
	VK_LAUNCH_APP2         = 183 // Windows 2000: Start Application 2 key
	VK_OEM_1               = 186 // Windows 2000: For the US standard keyboard, the ';:' key
	VK_OEM_PLUS            = 187 // Windows 2000: For any country/region, the '+' key
	VK_OEM_COMMA           = 188 // Windows 2000: For any country/region, the ',' key
	VK_OEM_MINUS           = 189 // Windows 2000: For any country/region, the '-' key
	VK_OEM_PERIOD          = 190 // Windows 2000: For any country/region, the '.' key
	VK_OEM_2               = 191 // Windows 2000: For the US standard keyboard, the '/?' key
	VK_OEM_3               = 192 // Windows 2000: For the US standard keyboard, the '`~' key
	VK_OEM_4               = 219 // Windows 2000: For the US standard keyboard, the '[{' key
	VK_OEM_5               = 220 // Windows 2000: For the US standard keyboard, the '\|' key
	VK_OEM_6               = 221 // Windows 2000: For the US standard keyboard, the ']}' key
	VK_OEM_7               = 222 // Windows 2000: For the US standard keyboard, the 'single-quote/double-quote' key
	VK_OEM_8               = 223 //
	VK_OEM_102             = 226 // Windows 2000: Either the angle bracket key or the backslash key on the RT 102-key keyboard
	VK_PROCESSKEY          = 229 // Windows 95/98, Windows NT 4.0, Windows 2000: IME PROCESS key
	VK_PACKET              = 231 // Windows 2000: Used to pass Unicode characters as if they were keystrokes. The VK_PACKET key is the low word of a 32-bit Virtual Key value used for non-keyboard input methods. For more information, see Remark in KEYBDINPUT, SendInput, WM_KEYDOWN, and WM_KEYUP
	VK_ATTN                = 246 // Attn key
	VK_CRSEL               = 247 // CrSel key
	VK_EXSEL               = 248 // ExSel key
	VK_EREOF               = 249 // Erase EOF key
	VK_PLAY                = 250 // Play key
	VK_ZOOM                = 251 // Zoom key
	VK_NONAME              = 252 // Reserved for future use
	VK_PA1                 = 253 // PA1 key
	VK_OEM_CLEAR           = 254 // Clear key

)

var id []int16
var user32 *syscall.DLL

type Hotkey struct {
	Modifiers int    // Mask of modifiers
	KeyCode   int    // Key code, e.g. 'A'
	CallBack  func() //
	ID        int16  // Unique id
}

type HotkeyEvents struct {
	Hotkeys []Hotkey
	stop    chan bool
	run     chan bool
	flag    bool

	ContinueID int16
}

func init() {
	user32 = syscall.MustLoadDLL("user32")
	id = make([]int16, 0)
}

func New() *HotkeyEvents {
	he := new(HotkeyEvents)
	he.stop = make(chan bool)
	he.run = make(chan bool)
	he.flag = false
	he.ContinueID = -1

	go func() {
		peekmsg := user32.MustFindProc("PeekMessageW")

		type MSG struct {
			HWND   uintptr
			UINT   uintptr
			WPARAM int16
			LPARAM int64
			DWORD  int32
			POINT  struct{ X, Y int64 }
		}

		for {
			select {
			case <-he.run:
				he.flag = true
			case <-he.stop:
				he.flag = false
			default:
				var msg = &MSG{}
				peekmsg.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0, 1)
				id := msg.WPARAM
				if (he.ContinueID == id) || (he.flag && id != 0) {
					for _, h := range he.Hotkeys {
						if h.ID == id {
							h.CallBack()
						}
					}
				}
			}
			time.Sleep(time.Millisecond * 50)
		}
	}()

	return he
}

func (he *HotkeyEvents) Bind(modi int, key int, callback func()) Hotkey {
	reghotkey := user32.MustFindProc("RegisterHotKey")
	h := Hotkey{
		Modifiers: modi,
		KeyCode:   key,
		CallBack:  callback,
	}
	h.ID = int16(len(id) + 1)
	id = append(id, h.ID)

	r1, _, err := reghotkey.Call(
		0, uintptr(h.ID), uintptr(h.Modifiers), uintptr(h.KeyCode))
	if r1 == 1 {
		he.Hotkeys = append(he.Hotkeys, h)
		return h
	} else {
		fmt.Println("Failed to register", h, ", error:", err)
	}
	return h
}

func (he *HotkeyEvents) BindContinue(modi int, key int, callback func()) Hotkey {
	h := he.Bind(modi, key, func() {
		go func() {
			he.run <- true
		}()
		callback()
	})
	he.ContinueID = h.ID
	return h
}

func (he *HotkeyEvents) Listen() {
	he.run <- true
}
func (he *HotkeyEvents) Stop() {
	go func() { he.stop <- true }()
}
