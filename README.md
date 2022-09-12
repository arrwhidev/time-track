# Time Track

Command line tool for time tracking backed by Google Sheets.

## Setup

- Create a [Google Cloud project](https://developers.google.com/workspace/guides/create-project) and enable the `Google Sheets API`.
- Create `Service account` [credentials](https://developers.google.com/workspace/guides/create-credentials).
- Create a new Google Sheet, click the Share button and share it with the Service account email address.
- Create a configuration file.
- Start tracking!

## Configuration

Manually create `/tmp/tt/config.json`

```
{
    "spreadsheetId": "",
    "email": "",
    "privateKey": "",
	"privateKeyId": "",
    "mappings": [
        "admin",
        "strategy",
        "operational"
    ]
}
```

## Google Sheet Configuration

- The sheeet should be named `Sheet1`
- Column `A` is reserved for the date, cell `A1` is the heading for the date -- you can name it anything.
- Row `1` should contain the headings, starting from cell `B1`, these should match the values in the JSON configuration file `mappings` attribute.g


## Usage

- **Add time** `tt add admin 20m` -- Add `20` minutes to the value of `admin`
- **Add time** `tt add admin 1h` -- Add `1` hour to the value of `admin`
- **Set time** `tt set admin 45m` -- Set the value for `admin` to `45` minutes
- **Sync** `tt sync` -- Sync today's data to your Google Sheet
- **Sync** `tt sync -F` -- Force Sync today's data to your Google Sheet

## Future Improvements

- Add support for multi-day sync (at the moment, the tool will only sync a single day).
- Remove requirement for configuration file to be handcrafted.