@set cPath=%cd:\=\\%
@set MODULE_ADDR=lab.xasea2.loc/root/go-grpc

@dir ..\ /s /b /a-d /o:s|findstr /v .git|findstr /v .idea|findstr /v tools|findstr /v swg|findstr /v README > .\web
@.\sed.exe -i -e "s#%cPath:\\tools=%#..#g" -e "s#\\#\\#g" .\web
@for /f %%i in (.\web) do @.\sed.exe -i "s#gitee.com/vipex/go-grpc#%MODULE_ADDR%#g" %%i

@rem PAUSE