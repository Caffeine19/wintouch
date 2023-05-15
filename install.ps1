$destination="$env:LOCALAPPDATA\wintouch\wintouch_windows_amd64.exe"
New-Item -ItemType Directory -Path "$env:LOCALAPPDATA\wintouch" -Force
Copy-Item -Path "./wintouch_windows_amd64.exe" -Destination $destination -Force

$newPath = $env:Path + ";$destination"

[System.Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")

New-Alias -Name touch -Value wintouch_windows_amd64.exe