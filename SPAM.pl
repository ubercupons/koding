#!/usr/bin/perl -w
# Hemp cURL
# Spammers mode ON
 
use strict;
use Net::SMTP::TLS;
 
die("$0 <user> <pass> <emails> <engenharia> <assunto>") if(@ARGV != 5);
my($user,$pass,$email,$eng,$assunto) = @ARGV;
my $smtp = new Net::SMTP::TLS(
 'smtp.gmail.com',
 Port    => 587,
 User    => $user,
 Password=> $pass,
 Timeout => 30
) || die($!);
 
open(EMAILS,'<'.$email) || die($!);
my @emails = <emails>;
close(EMAILS);
open(ENG,'<'.$eng) || die($!);
my @e = <eng>;
my $en = join('',@e);
close(ENG);
 
foreach(@emails){
 chomp($_);
 $smtp->mail($user);
 $smtp->recipient($_);
 $smtp->data();
 $smtp->datasend("To: $_\n");
 $smtp->datasend("From: Nome de quem mandou <$user>\n");
 $smtp->datasend("Content-Type: text/html \n");
 $smtp->datasend("Subject: $assunto");
 $smtp->datasend("\n");
 $smtp->datasend("$en");
 $smtp->datasend("\n");
 $smtp->dataend();
 print "Enviado para $_\n";
}
 
$smtp->quit;
H3mp coding...
