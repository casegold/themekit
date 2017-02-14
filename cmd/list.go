package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/Shopify/themekit/kit"
)

var listCmd = &cobra.Command{
	Use:   "list <filenames>",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		setFlagConfig()
		themes, err := kit.Themes(configPath, environment)

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
		fmt.Fprintln(w, strings.Join([]string{"ID", "Name", "Role", "Previewable", "Processing"}, "\t "))
		fmt.Fprintln(w, strings.Join([]string{"----------", "----------", "----------", "----------", "----------"}, "\t "))
		for _, theme := range themes {
			fmt.Fprintln(w,
				strings.Join([]string{
					strconv.Itoa(int(theme.ID)),
					theme.Name,
					theme.Role,
					strconv.FormatBool(theme.Previewable),
					strconv.FormatBool(theme.Processing),
				}, "\t "),
			)
		}
		w.Flush()

		return err
	},
}
