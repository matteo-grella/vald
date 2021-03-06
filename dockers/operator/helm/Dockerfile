#
# Copyright (C) 2019-2020 Vdaas.org Vald team ( kpango, rinx, kmrmt )
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
ARG OPERATOR_SDK_VERSION
ARG VERSION="v1alpha1"
ARG GROUP="vald.vdaas.org"
ARG VALD_KIND="ValdRelease"
ARG VALD_HELM_OPERATOR_KIND="ValdHelmOperatorRelease"

FROM vdaas/vald-base:latest AS builder
ARG OPERATOR_SDK_VERSION
ARG VERSION
ARG GROUP
ARG VALD_KIND
ARG VALD_HELM_OPERATOR_KIND
RUN { \
        echo "---"; \
        echo "- version: ${VERSION}"; \
        echo "  group: ${GROUP}"; \
        echo "  kind: ${VALD_KIND}"; \
        echo "  chart: helm-charts/vald"; \
        echo "- version: ${VERSION}"; \
        echo "  group: ${GROUP}"; \
        echo "  kind: ${VALD_HELM_OPERATOR_KIND}"; \
        echo "  chart: helm-charts/vald-helm-operator"; \
    } > /tmp/watches.yaml

FROM quay.io/operator-framework/helm-operator:${OPERATOR_SDK_VERSION}
LABEL maintainer "Vald team <vald@vdaas.org>"

COPY charts/vald ${HOME}/helm-charts/vald
COPY charts/vald-helm-operator ${HOME}/helm-charts/vald-helm-operator
COPY --from=builder /tmp/watches.yaml ${HOME}/watches.yaml
