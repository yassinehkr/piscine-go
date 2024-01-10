find -type f -name "*.sh" | cut -c 3- | sed "s/\.sh//g" | rev | cut -d '/' -f1 | rev
