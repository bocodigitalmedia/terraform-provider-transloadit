variable "s3_bucket_name" {
  type = "string"
}

variable "transloadit_auth_key" {
  type = "string"
}

variable "transloadit_auth_secret" {
  type = "string"
}

variable "aws_access_key_id" {
  type = "string"
}

variable "aws_secret_access_key" {
  type = "string"
}

variable "url_prefix" {
  type = "string"
}

provider "transloadit" {
  auth_key    = "${var.transloadit_auth_key}"
  auth_secret = "${var.transloadit_auth_secret}"
}

resource "transloadit_template" "test" {
  name = "audio-import.TEST"

  step {
    name = "imported"

    params {
      path   = "$${fields.uploadKey}"
      robot  = "/s3/import"
      key    = "${var.aws_access_key_id}"
      secret = "${var.aws_secret_access_key}"
      bucket = "${var.s3_bucket_name}"
    }
  }

  step {
    name = "mp3_standard"

    params {
      use    = "imported"
      preset = "mp3"
      result = true
      robot  = "/audio/encode"
    }
  }

  step {
    name = "mp3_standard_export"

    params {
      use        = "mp3_standard"
      robot      = "/s3/store"
      key        = "${var.aws_access_key_id}"
      acl        = "private"
      secret     = "${var.aws_secret_access_key}"
      bucket     = "boco-delphi-v2-development"
      path       = "resources/$${fields.resourceId}/$${assembly.id}/$${previous_step.name}/$${file.url_name}"
      url_prefix = "${var.url_prefix}"
    }
  }

  step {
    name = "mp3_standard_waveform"

    params {
      use              = "mp3_standard"
      format           = "image"
      width            = 200
      height           = 200
      background_color = "FFFFFFFF"
      center_color     = "333333FF"
      outer_color      = "333333FF"
      result           = false
      robot            = "/audio/waveform"
    }
  }

  step {
    name = "thumbnail_200x200"

    params {
      use    = "mp3_standard_waveform"
      result = true
      robot  = "/image/optimize"
    }
  }

  step {
    name = "thumbnail_200x200_export"

    params {
      use        = "thumbnail_200x200"
      robot      = "/s3/store"
      key        = "${var.aws_access_key_id}"
      acl        = "private"
      secret     = "${var.aws_secret_access_key}"
      bucket     = "${var.s3_bucket_name}"
      path       = "resources/$${fields.resourceId}/$${assembly.id}/$${previous_step.name}/$${file.url_name}"
      url_prefix = "${var.url_prefix}"
    }
  }
}
