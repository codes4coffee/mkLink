#Make Link
A symbolic link creator for github pages in golang
===
This simple app allows you to create symbolic URLs to other webpages that have longer or harder to read URLs. The way this app works is by creating a folder in your webroot with the name of the symbolic URL. For example, if I wanted foo.bar/git to point to github.com, the tool would create a folder called git and in that folder there would be an index.html page that redirects to github.com.

##Argument structure
mkLink <Symbolic name> <Actual URL excluding https://>