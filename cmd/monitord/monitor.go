package main

import (
    "os"
    "fmt"
    "github.com/spf13/cobra"
)

func NewCommand(RootCmd *cobra.Command) *cobra.Command {
    HelpCommand := &cobra.Command{
        Use: "help",
        Short: "A high performance webserver",
        Long: `Hugo provides its own webserver which builds and serves the site.
While hugo server is high performance, it is a webserver with limited options.
Many run it in production, but the standard behavior is for people to use it
in development and use a more full featured server such as Nginx or Caddy.
'hugo server' will avoid writing the rendered and served content to disk,
preferring to store it in memory.
By default hugo will also watch your files for any changes you make and
automatically rebuild the site. It will then live reload any open browser pages
and push the latest content to them. As most Hugo sites are built in a fraction
of a second, you will be able to save and see your changes nearly instantly.`,
    }
    
    VersionCommand := &cobra.Command{
        Use: "version",
        Short: "A high performance webserver",
        Long: `Hugo provides its own webserver which builds and serves the site.
While hugo server is high performance, it is a webserver with limited options.
Many run it in production, but the standard behavior is for people to use it
in development and use a more full featured server such as Nginx or Caddy.
'hugo server' will avoid writing the rendered and served content to disk,
preferring to store it in memory.
By default hugo will also watch your files for any changes you make and
automatically rebuild the site. It will then live reload any open browser pages
and push the latest content to them. As most Hugo sites are built in a fraction
of a second, you will be able to save and see your changes nearly instantly.`,
    }
    
    RootCmd.AddCommand(HelpCommand)
    RootCmd.AddCommand(VersionCommand)
    return RootCmd
}

func main() {
    
    RootCmd := &cobra.Command{
        Use: "monitord",
        Short: "Linux server status monitor",
        Run: func(cmd *cobra.Command, args []string) {
            // TODO ...
        },
    }
    
    // 添加子命令
    NewCommand(RootCmd)
    
    // TODO flags
    //Flags := RootCmd.Flags()
    //Flags.StringVar(&version, "version", "1.0.0", "this monitor version")
    RootCmd.Flags().Bool("cleanDestinationDir", false, "Remove files from destination not found in static directories")
    RootCmd.Flags().BoolP("buildDrafts", "D", false, "include content marked as draft")
    RootCmd.Flags().BoolP("buildFuture", "F", false, "include content with publishdate in the future")
    
    if err := RootCmd.Execute(); err != nil {
        fmt.Printf("%v", err)
        os.Exit(1)
    }
    
}