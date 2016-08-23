// Copyright © 2016 Phillip Dudley <Predatorian3@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"math/rand"
	"os"
    "strings"
    "strconv"
    "time"
	"github.com/spf13/cobra"
)

func checkerr( err error ) {
	if err != nil {
    	fmt.Println(err)
    	os.Exit(2)

    }
}

func roll_dice( dice string ) {
	var dice_result int = 0
    // This will be an array of dice
    dice_array := strings.Split( dice, "d" )
    // use strconv to convert arrays indexes to Integers
    dice_num, err := strconv.Atoi( dice_array[0] )
    // Check for errors.
    checkerr(err)
    dice_type, err := strconv.Atoi( dice_array[1] )
    checkerr(err)
    // loop through the number of dice provided
    for i := 1; i <= dice_num; i++ {
    	// Generate a seed to get a truely random number.
    	// otherwise it'll be the same number every build.
    	seed_number := rand.New(rand.NewSource(time.Now().UnixNano()))
    	// Use the seed to generate an Integer to the maximum of what the
    	// dice type is. 
	    number_result := seed_number.Intn(dice_type)
	    // Inciment by 1 so that there are no zeros.
	    number_result = number_result + 1
	    // Display the rolls. 
	    fmt.Println( "Roll #",i,":", number_result ) 
	    // Keep a total of the dice numbers. 
	    dice_result = dice_result + number_result 
    }
    fmt.Println( "Total:", dice_result )
}

// rollCmd represents the roll command
var rollCmd = &cobra.Command{
	Use:   "roll #d#",
	Short: "Roll a set of dice specified",
	Long: `Roll a set of dice specified by the number

(Number of dice)d(number of sides)

This denotes the number and type of dice being rolled.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("roll called")
		roll_dice(args[0])
	},
}

func init() {
	RootCmd.AddCommand(rollCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rollCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rollCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
