package tyk

var apiTemplates = `
{{ define "default" }}
{
    "name": "{{.Name}}{{ range $i, $e := .GatewayTags }} #{{$e}}{{ end }}",
	"slug": "{{.Slug}}",
    "org_id": "{{.Org}}",
    "use_keyless": true,
    "definition": {
        "location": "header",
        "key": "x-api-version",
        "strip_path": true
    },
    "version_data": {
        "not_versioned": true,
        "versions": {
            "Default": {
                "name": "Default",
                "use_extended_paths": true,
				"global_headers": {
                    "X-Tyk-Ingress-Request-ID": "$tyk_context.request_id"
                },
				"paths": {
                    "ignored": [],
                    "white_list": [],
                    "black_list": []
                }
            }
        }
    },
    "proxy": {
        "listen_path": "{{.ListenPath}}",
        "target_url": "{{.Target}}",
        "strip_listen_path": true
    },
	"domain": "{{.HostName}}",
	"response_processors": [],
	 "custom_middleware": {
        "pre": [],
        "post": [],
        "post_key_auth": [],
        "auth_check": {
            "name": "",
            "path": "",
            "require_session": false
        },
        "response": [],
        "driver": "",
        "id_extractor": {
            "extract_from": "",
            "extract_with": "",
            "extractor_config": {}
        }
    },
	"config_data": {},
	"allowed_ips": [],
    "disable_rate_limit": true,
    "disable_quota": true,
    "cache_options": {
        "cache_timeout": 60,
        "enable_cache": true
    },
    "active": true,
    "tags": [{{ range $i, $e := .GatewayTags }}{{ if $i }},{{ end }}"{{ $e }}"{{ end }}],
    "enable_context_vars": true,
	"certificates": [{{ range $i, $e := .CertificateID }}{{ if $i }},{{ end }}"{{ $e }}"{{ end }}]
}
{{ end }}

{{ define "default-mesh" }}
{
    "name": "{{.Name}}{{ range $i, $e := .GatewayTags }} #{{$e}}{{ end }}",
	"slug": "{{.Slug}}",
    "org_id": "{{.Org}}",
    "use_keyless": true,
    "definition": {
        "location": "header",
        "key": "x-api-version",
        "strip_path": true
    },
    "version_data": {
        "not_versioned": true,
        "versions": {
            "Default": {
                "name": "Default",
                "use_extended_paths": true,
				"global_headers": {
                    "X-Tyk-Mesh-Request-ID": "$tyk_context.request_id"
                },
				"paths": {
                    "ignored": [],
                    "white_list": [],
                    "black_list": []
                }
            }
        }
    },
    "proxy": {
        "listen_path": "{{.ListenPath}}",
        "target_url": "{{.Target}}",
        "strip_listen_path": true
    },
	"domain": "{{.HostName}}",
	"response_processors": [],
	 "custom_middleware": {
        "pre": [],
        "post": [],
        "post_key_auth": [],
        "auth_check": {
            "name": "",
            "path": "",
            "require_session": false
        },
        "response": [],
        "driver": "",
        "id_extractor": {
            "extract_from": "",
            "extract_with": "",
            "extractor_config": {}
        }
    },
	"config_data": {},
	"allowed_ips": [],
    "disable_rate_limit": true,
    "disable_quota": true,
    "cache_options": {
        "cache_timeout": 60,
        "enable_cache": true
    },
    "active": true,
    "tags": [{{ range $i, $e := .GatewayTags }}{{ if $i }},{{ end }}"{{ $e }}"{{ end }}],
    "enable_context_vars": true,
	"certificates": [{{ range $i, $e := .CertificateID }}{{ if $i }},{{ end }}"{{ $e }}"{{ end }}]
}
{{ end }}

{{ define "default-inbound" }}
{
    "name": "{{.Name}}{{ range $i, $e := .GatewayTags }} #{{$e}}{{ end }}",
	"slug": "{{.Slug}}",
    "org_id": "{{.Org}}",
    "use_keyless": true,
    "definition": {
        "location": "header",
        "key": "x-api-version",
        "strip_path": true
    },
    "version_data": {
        "not_versioned": true,
        "versions": {
            "Default": {
                "name": "Default",
                "use_extended_paths": true,
				"global_headers": {
                    "X-Tyk-Mesh-Inbound-Request-ID": "$tyk_context.request_id"
                },
				"paths": {
                    "ignored": [],
                    "white_list": [],
                    "black_list": []
                }
            }
        }
    },
    "proxy": {
        "listen_path": "{{.ListenPath}}",
        "target_url": "{{.Target}}",
        "strip_listen_path": true
    },
	"domain": "{{.HostName}}",
	"response_processors": [],
	 "custom_middleware": {
        "pre": [],
        "post": [],
        "post_key_auth": [],
        "auth_check": {
            "name": "",
            "path": "",
            "require_session": false
        },
        "response": [],
        "driver": "",
        "id_extractor": {
            "extract_from": "",
            "extract_with": "",
            "extractor_config": {}
        }
    },
	"config_data": {},
	"allowed_ips": [],
    "disable_rate_limit": true,
    "disable_quota": true,
    "cache_options": {
        "cache_timeout": 60,
        "enable_cache": true
    },
    "active": true,
    "tags": [{{ range $i, $e := .GatewayTags }}{{ if $i }},{{ end }}"{{ $e }}"{{ end }}],
    "enable_context_vars": true,
	"certificates": [{{ range $i, $e := .CertificateID }}{{ if $i }},{{ end }}"{{ $e }}"{{ end }}]
}
{{ end }}

`
