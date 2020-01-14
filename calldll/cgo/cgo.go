package main

/*
#include <Windows.h>
#include <winuser.h>
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}

int DisplayResourceNAMessageBox()
{
    int msgboxID = MessageBoxW(
        NULL,
        (LPCWSTR)L"中文This test is Done.",
        (LPCWSTR)L"Done Title",
        MB_ICONWARNING | MB_CANCELTRYCONTINUE | MB_DEFBUTTON2
    );

    switch (msgboxID)
    {
    case IDCANCEL:
		printf("printint: %d\n", msgboxID);
        break;
    case IDTRYAGAIN:
		printf("printint: %d\n", msgboxID);
        break;
    case IDCONTINUE:
		printf("printint: %d\n", msgboxID);
        break;
    }
    return msgboxID;
}
*/
import "C"

func main() {
	//v := 42
	//C.printint(C.int(v))
	C.DisplayResourceNAMessageBox()
}
