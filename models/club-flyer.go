package models

// ChessClubFlyer ... Holds the Respons Data From Creating a new Club Flyer
type ChessClubFlyer struct {
	Year            int64
	Name            string
	Content         string
	LeftImage       image
	RightImage      image
	Discord         link
	Slack           link
	LiChessTeam     link
	ChessdotcomTeam link
	Twitch          link
	Instagram       link
	Constitution    link
	Roster          link
}

// FillDefaults ... Constructor Function: Setting default values: if no values present
func (ChessClubFlyer *ChessClubFlyer) FillDefaults() {

	// Year, Name, Content
	if ChessClubFlyer.Year <= 2021 {ChessClubFlyer.Year = 2021}
	if ChessClubFlyer.Name == "" {ChessClubFlyer.Name = "MakeSchool_Chess_Club_Flyer"}
	if ChessClubFlyer.Content == "" {ChessClubFlyer.Content = "Test Info"}

	// Left Image
	if ChessClubFlyer.LeftImage.SRC == "" {ChessClubFlyer.LeftImage.SRC = "2d_chess_board.png"}
	if ChessClubFlyer.LeftImage.ALT == "" {ChessClubFlyer.LeftImage.ALT = "2D Chess Board"}

	// Right Image
	if ChessClubFlyer.RightImage.SRC == "" {ChessClubFlyer.RightImage.SRC = "3d_chess_board.png"}
	if ChessClubFlyer.RightImage.ALT == "" {ChessClubFlyer.RightImage.ALT = "3D Chess Board"}


	// Chess Club Constitution
	if ChessClubFlyer.Constitution.Display == "" {ChessClubFlyer.Constitution.Display = "MakeSchool Chess Club Constitution"}
	if ChessClubFlyer.Constitution.URL == "" {ChessClubFlyer.Constitution.URL = "https://docs.google.com/document/d/10ZfD5Nqg3FtCFSnnLLMcR9lPokXMMG8WOIRWl_tAHbo/edit?usp=sharing"}

	// Chess Club Roster
	if ChessClubFlyer.Roster.Display == "" {ChessClubFlyer.Roster.Display = "MakeSchool Chess Club Roster"}
	if ChessClubFlyer.Roster.URL == "" {ChessClubFlyer.Roster.URL = "https://docs.google.com/spreadsheets/d/1KMwY6cZmziu6rvD3e-ErarretxMFtCwj79X1sRlvUoQ/edit?usp=sharing"}


	// Chat Channels

	// Discord
	if ChessClubFlyer.Discord.Display == "" {ChessClubFlyer.Discord.Display = "discord.gg/MAR5NJYpfv"}
	if ChessClubFlyer.Discord.URL == "" {ChessClubFlyer.Discord.URL = "https://discord.gg/MAR5NJYpfv"}

	// Slack
	if ChessClubFlyer.Slack.Display == "" {ChessClubFlyer.Slack.Display = "#chess-club"}
	if ChessClubFlyer.Slack.URL == "" {ChessClubFlyer.Slack.URL = "https://makeschool.us20.list-manage.com/track/click?u=b6d3ee0bfdf3579e1533596e6&id=aa8c18e854&e=cff62999ba"}


	// Chess Platforms

	// lichess.org
	if ChessClubFlyer.LiChessTeam.Display == "" {ChessClubFlyer.LiChessTeam.Display = "MakeSchool Team on lichess.org"}
	if ChessClubFlyer.LiChessTeam.URL == "" {ChessClubFlyer.LiChessTeam.URL = "https://lichess.org/team/makeschool"}

	// chess.com
	if ChessClubFlyer.ChessdotcomTeam.Display == "" {ChessClubFlyer.ChessdotcomTeam.Display = "MakeSchool Team on Chess.com"}
	if ChessClubFlyer.ChessdotcomTeam.URL == "" {ChessClubFlyer.ChessdotcomTeam.URL = "https://www.chess.com/club/make-school/join"}


	// Make School Links
	if ChessClubFlyer.Twitch.Display == "" {ChessClubFlyer.Twitch.Display = "MakeSchool on Twitch.tv"}
	if ChessClubFlyer.Twitch.URL == "" {ChessClubFlyer.Twitch.URL = "twitch.tv/make_school"}

	if ChessClubFlyer.Instagram.Display == "" {ChessClubFlyer.Instagram.Display = "MakeSchool on Instagram.com"}
	if ChessClubFlyer.Instagram.URL == "" {ChessClubFlyer.Instagram.URL = "https://Instagram.com/make_school"}
}