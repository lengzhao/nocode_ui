package widget

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func init() {
	Regist("icon", icon{})
}

type icon struct {
}

var iconNames []fyne.ThemeIconName = []fyne.ThemeIconName{
	theme.IconNameCancel,
	theme.IconNameConfirm,
	theme.IconNameDelete,
	theme.IconNameSearch,
	theme.IconNameSearchReplace,
	theme.IconNameMenu,
	theme.IconNameMenuExpand,
	theme.IconNameCheckButton,
	theme.IconNameCheckButtonChecked,
	theme.IconNameRadioButton,
	theme.IconNameRadioButtonChecked,

	theme.IconNameContentAdd,
	theme.IconNameContentClear,
	theme.IconNameContentRemove,
	theme.IconNameContentCut,
	theme.IconNameContentCopy,
	theme.IconNameContentPaste,
	theme.IconNameContentRedo,
	theme.IconNameContentUndo,

	theme.IconNameColorAchromatic,
	theme.IconNameColorChromatic,
	theme.IconNameColorPalette,
	theme.IconNameDocument,
	theme.IconNameDocumentCreate,
	theme.IconNameDocumentPrint,
	theme.IconNameDocumentSave,

	theme.IconNameMoreHorizontal,
	theme.IconNameMoreVertical,

	theme.IconNameInfo,
	theme.IconNameQuestion,
	theme.IconNameWarning,
	theme.IconNameError,

	theme.IconNameMailAttachment,
	theme.IconNameMailCompose,
	theme.IconNameMailForward,
	theme.IconNameMailReply,
	theme.IconNameMailReplyAll,
	theme.IconNameMailSend,

	theme.IconNameMediaMusic,
	theme.IconNameMediaPhoto,
	theme.IconNameMediaVideo,
	theme.IconNameMediaFastForward,
	theme.IconNameMediaFastRewind,
	theme.IconNameMediaPause,
	theme.IconNameMediaPlay,
	theme.IconNameMediaRecord,
	theme.IconNameMediaReplay,
	theme.IconNameMediaSkipNext,
	theme.IconNameMediaSkipPrevious,
	theme.IconNameMediaStop,

	theme.IconNameNavigateBack,
	theme.IconNameMoveDown,
	theme.IconNameNavigateNext,
	theme.IconNameMoveUp,
	theme.IconNameArrowDropDown,
	theme.IconNameArrowDropUp,

	theme.IconNameFile,
	theme.IconNameFileApplication,
	theme.IconNameFileAudio,
	theme.IconNameFileImage,
	theme.IconNameFileText,
	theme.IconNameFileVideo,
	theme.IconNameFolder,
	theme.IconNameFolderNew,
	theme.IconNameFolderOpen,
	theme.IconNameHelp,
	theme.IconNameHistory,
	theme.IconNameHome,
	theme.IconNameSettings,
	theme.IconNameViewFullScreen,
	theme.IconNameViewRefresh,
	theme.IconNameViewRestore,
	theme.IconNameViewZoomFit,
	theme.IconNameViewZoomIn,
	theme.IconNameViewZoomOut,

	theme.IconNameVisibility,
	theme.IconNameVisibilityOff,

	theme.IconNameVolumeDown,
	theme.IconNameVolumeMute,
	theme.IconNameVolumeUp,

	theme.IconNameDownload,
	theme.IconNameComputer,
	theme.IconNameStorage,
	theme.IconNameUpload,

	theme.IconNameAccount,
	theme.IconNameLogin,
	theme.IconNameLogout,

	theme.IconNameList,
	theme.IconNameGrid,
}

func (i icon) Create(conf Config, childs []fyne.CanvasObject) (fyne.CanvasObject, error) {
	icon := fyne.CurrentApp().Settings().Theme().Icon(fyne.ThemeIconName(conf.Name))

	w := widget.NewIcon(icon)
	if conf.Width > 1 && conf.Height > 1 {
		w.Resize(fyne.NewSize(conf.Width, conf.Height))
	}
	return w, nil
}

func (i icon) Help() string {
	out := fmt.Sprintln("type: icon")
	out += fmt.Sprintln(" describtion: new icon.")
	out += fmt.Sprintln(" Name: the icon name.")
	out += fmt.Sprintln(" icon name list:")
	for i, it := range iconNames {
		out += fmt.Sprintln(" ", i, ". ", it)
	}
	return out
}
