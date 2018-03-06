<?php
    $ID_user     = "jesuisunuserID";
    $URL_back    = "http://hunt.ly:5050";
    echo "$URL_back/token?id=$ID_user";
    print file_get_contents("$URL_back/token?id=$ID_user");

?>