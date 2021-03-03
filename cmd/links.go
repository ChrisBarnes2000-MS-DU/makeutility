package cmd

import (
	"github.com/foize/go.sgr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(linksCommand)
	linksCommand.AddCommand(lichessLinks)
	linksCommand.AddCommand(contactLinks)
	linksCommand.AddCommand(otherLinks)
}

var linksCommand = &cobra.Command{
	Use:   "links",
	Short: "[Commands]: contact, lichess, all | List links related to chess",
	Long:  "List links related to Make School's Chess Club",
}

var lichessLinks = &cobra.Command{
	Use:   "lichess",
	Short: "List links related to lichess.org & Make School's Chess Club",
	Long:  "List links related to lichess.org & Make School's Chess Club",
	Run: func(cmd *cobra.Command, args []string) {
		sgr.Println("https://lichess.org/ads")
		sgr.Println("https://lichess.org/lag")
		sgr.Println("https://lichess.org/faq")
		sgr.Println("https://lichess.org/api")
		sgr.Println("----------------------------------")
		sgr.Println("https://lichess.org/blog/V0KrLSkAAMo3hsi4/study-chess-the-lichess-way")
		sgr.Println("https://lichess.org/analysis")
		sgr.Println("https://lichess.org/learn")
		sgr.Println("https://lichess.org/practice")
		sgr.Println("----------------------------------")
	},
}

var contactLinks = &cobra.Command{
	Use:   "contact",
	Short: "Related Links To Contact MS Chess Club",
	Long:  "Links Related To Contacting Make School's Chess Club",
	Run: func(cmd *cobra.Command, args []string) {
		sgr.Println("https://discord.gg/MAR5NJYpfv")
		sgr.Println("#chess-club")
		sgr.Println("https://lichess.org/team/makeschool")
		sgr.Println("https://chess.com/club/make-school/join")
		sgr.Println("https://twitch.tv/make_school")
		sgr.Println("https://instagram.com/make_school")
		sgr.Println("----------------------------------")
		sgr.Println("https://lichess.org/@/cbarnes2000/")
		sgr.Println("https://lichess.org/@/cbarnes2000/follow")
		sgr.Println("https://lichess.org/@/cbarnes2000/tournaments/created")
		sgr.Println("----------------------------------")
	},
}

var otherLinks = &cobra.Command{
	Use:   "other",
	Short: "List links related to chess",
	Long:  "List links related to Chess Club",
	Run: func(cmd *cobra.Command, args []string) {
		sgr.Println("https://ecf.octoknight.com/")
		sgr.Println("----------------------------------")
		sgr.Println("https://www.englishchess.org.uk/ecf-membership-partners-and-benefits/")
		sgr.Println("https://www.englishchess.org.uk/ecf-membership-rates-and-joining-details/")
		sgr.Println("----------------------------------")
		sgr.Println("https://rotherhamonlinechess.azurewebsites.net/tournaments")
		sgr.Println("----------------------------------")
		sgr.Println("https://www.englishchess.org.uk/")
		sgr.Println("http://ecfgrading.org.uk/new/menu.php")
		sgr.Println("----------------------------------")
		sgr.Println("https://new.uschess.org/")
		sgr.Println("http://www.uschess.org/index.php/Players-Ratings/")
		sgr.Println("----------------------------------")
		sgr.Println("https://fide.com/")
		sgr.Println("https://ratings.fide.com/")
		sgr.Println("----------------------------------")
		sgr.Println("https://www.regencychess.co.uk/blog/2012/11/chess-noob-20-ratings/")
		sgr.Println("----------------------------------")
		sgr.Println("https://www.milibrary.org/chess")
		sgr.Println("https://www.milibrary.org/blog/extended-closure-faq")
		sgr.Println("https://www.milibrary.org/chess-links")
		sgr.Println("----------------------------------")
		sgr.Println("https://www.masterclass.com/classes/garry-kasparov-teaches-chess")
		sgr.Println("----------------------------------")
	},
}
