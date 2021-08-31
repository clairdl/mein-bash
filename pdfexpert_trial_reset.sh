#!/bin/bash

# --------------------------------------------------------
# this script will remove config files for PDF Expert,
# allowing you to indefinitely claim the 7-day free trial
# --------------------------------------------------------
# If you want to schedule this to run every 7 days, see:
# https://crontab.guru/every-7-days
# --------------------------------------------------------

plist=~/Library/Preferences/com.readdle.PDFExpert-Mac.plist
supp0=~/Library/Application\ Support/com.readdle.PDFExpert-Mac
supp1=~/Library/Application\ Support/PDF\ Expert

if (ps aux | grep "PDF Expert" | grep -v "grep" > /dev/null)
  then
    echo "Please close PDF Expert (pid: ${$(pgrep "PDF Expert")}) before running this script"
  fi

if [[ -f $plist ]] && [[ -d $supp0 ]] && [[ -d $supp1 ]]
  then
    rm $plist
    rm -rf $supp0
    rm -rf $supp1
    echo "Free trial reset! \nTip: you can provide a fake email address when PDF Expert prompts you \ni.e. fake@mail.com will work! :) "
  else
    echo "PDF Expert config files have moved ergo this script no longer works. \nTweet me @eulalia0495 and I'll fix it."
  fi
