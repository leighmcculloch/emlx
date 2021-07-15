# emlx
Email file (.eml) attachment and embedded file extractor.

## Install

```
brew install 4d63/emlx/emlx
```
or

```
go install github.com/leighmcculloch/emlx@latest
```

## Usage

```
emlx [emlfile.eml] [emlfile.eml] ...
```

Directories are created with the eml files date and subject.

Attachments and embedded files are written to the directories.

Files with duplicate names overwrite and is not handled well.

