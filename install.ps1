$destination="$env:LOCALAPPDATA\wintouch\wintouch_windows_amd64.exe"
New-Item -ItemType Directory -Path "$env:LOCALAPPDATA\wintouch" -Force
Copy-Item -Path "./wintouch_windows_amd64.exe" -Destination $destination -Force

$newPath = $envPath + ";$destination"
[System.Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")

 NEw-Alias -Name touch -Value wintouch_windows_amd64.exe