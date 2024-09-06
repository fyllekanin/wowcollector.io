<script lang="ts" setup>
import { Icons } from '~/constants';

definePageMeta({
  layout: 'empty',
});

const router = useRouter();
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 py-8">
    <UButton
      class="mb-4"
      variant="ghost"
      :icon="Icons.BACK"
      @click="router.back()"
    >
      Go back
    </UButton>

    <h1 class="text-3xl font-bold mb-6">Terms and Conditions</h1>
    <p class="text-sm text-gray-500 mb-8">Last updated: September 02, 2024</p>

    <p class="mb-6">
      Welcome to WoW Collector (the "Website"). By accessing or using our
      Website, you agree to comply with and be bound by the following Terms and
      Conditions (the "Terms"). If you do not agree to these Terms, you should
      not use this Website. These Terms constitute a legal agreement between you
      (the "User") and WoW Collector ("we," "us," or "our").
    </p>

    <h2 class="text-2xl font-semibold mb-4">1. Acceptance of Terms</h2>
    <p class="mb-6">
      By accessing or using the Website, you acknowledge that you have read,
      understood, and agree to be bound by these Terms and our Privacy Policy,
      which is incorporated herein by reference. If you do not agree to these
      Terms, you must immediately stop using the Website.
    </p>

    <h2 class="text-2xl font-semibold mb-4">2. Eligibility</h2>
    <p class="mb-6">
      To use the Website, you must be at least 13 years old. By using the
      Website, you represent and warrant that you meet this eligibility
      requirement. If you are using the Website on behalf of a company,
      organization, or other legal entity, you represent and warrant that you
      have the authority to bind that entity to these Terms.
    </p>

    <h2 class="text-2xl font-semibold mb-4">3. User Accounts and Sign-In</h2>
    <p class="mb-6">
      Users may sign in to the Website using third-party providers such as
      Discord or Battle.net. By using third-party sign-in options, you agree to
      comply with the terms and policies of those providers. You are responsible
      for maintaining the confidentiality of your account credentials and for
      all activities that occur under your account. You must immediately notify
      us of any unauthorized use of your account.
    </p>

    <h2 class="text-2xl font-semibold mb-4">4. User-Generated Content</h2>
    <p class="mb-6">
      Users may create and/or upload content to the Website, including but not
      limited to text, images, and other media ("User Content"). By uploading or
      creating User Content, you grant us a non-exclusive, worldwide,
      royalty-free, perpetual, irrevocable license to use, reproduce, modify,
      adapt, publish, translate, distribute, display, and create derivative
      works from such User Content in any form, media, or technology, whether
      now known or hereafter developed.
    </p>
    <p class="mb-6">
      You represent and warrant that you own or have the necessary licenses,
      rights, consents, and permissions to grant the above license to us. You
      are solely responsible for your User Content and any consequences of
      posting or publishing it.
    </p>

    <h2 class="text-2xl font-semibold mb-4">5. Prohibited Conduct</h2>
    <ul class="list-disc pl-6 mb-8 space-y-2">
      <li>
        Post or transmit any content that is unlawful, harmful, threatening,
        abusive, harassing, defamatory, vulgar, obscene, or otherwise
        objectionable.
      </li>
      <li>
        Engage in any conduct that would violate any applicable law, regulation,
        or third-party rights.
      </li>
      <li>
        Upload viruses, malware, or other malicious code that could harm the
        functionality of the Website or the experience of any other user.
      </li>
      <li>
        Attempt to gain unauthorized access to any portion or feature of the
        Website, or any other systems or networks connected to the Website.
      </li>
    </ul>

    <h2 class="text-2xl font-semibold mb-4">6. Intellectual Property Rights</h2>
    <p class="mb-6">
      All content, trademarks, service marks, logos, and visual elements
      displayed on the Website (collectively, "Our Content") are our exclusive
      property. You are not permitted to use Our Content without our prior
      written consent. Nothing in these Terms grants you any right, title, or
      interest in Our Content.
    </p>

    <h2 class="text-2xl font-semibold mb-4">7. Feedback and Suggestions</h2>
    <p class="mb-6">
      We welcome your feedback, suggestions, and ideas for improving the Website
      ("Feedback"). If you provide us with Feedback, you agree that we are free
      to use, disclose, reproduce, modify, license, distribute, and otherwise
      exploit the Feedback without any obligation to you, including any payment
      or compensation of any kind.
    </p>

    <h2 class="text-2xl font-semibold mb-4">
      8. Third-Party Links and Services
    </h2>
    <p class="mb-6">
      The Website may contain links to third-party websites, services, or
      resources that are not owned or controlled by us. We have no control over,
      and assume no responsibility for, the content, privacy policies, or
      practices of any third-party websites or services. You acknowledge and
      agree that we are not responsible or liable, directly or indirectly, for
      any damage or loss caused or alleged to be caused by or in connection with
      the use of any such content, goods, or services available on or through
      any such third-party websites or services.
    </p>

    <h2 class="text-2xl font-semibold mb-4">9. Termination</h2>
    <p class="mb-6">
      We reserve the right to terminate or suspend your access to the Website,
      with or without notice, for any reason, including but not limited to your
      breach of these Terms. Upon termination, all licenses and other rights
      granted to you under these Terms will immediately cease.
    </p>

    <h2 class="text-2xl font-semibold mb-4">10. Disclaimers</h2>
    <p class="mb-6">
      The Website is provided on an "AS IS" and "AS AVAILABLE" basis. We make no
      representations or warranties of any kind, express or implied, as to the
      operation of the Website, or the information, content, materials, or
      products included on the Website. You expressly agree that your use of the
      Website is at your sole risk.
    </p>

    <h2 class="text-2xl font-semibold mb-4">11. Limitation of Liability</h2>
    <p class="mb-6">
      To the fullest extent permitted by law, we shall not be liable for any
      direct, indirect, incidental, special, consequential, or punitive damages,
      including but not limited to, damages for loss of profits, goodwill, use,
      data, or other intangible losses arising out of or in connection with your
      use of the Website.
    </p>

    <h2 class="text-2xl font-semibold mb-4">12. Indemnification</h2>
    <p class="mb-6">
      You agree to indemnify, defend, and hold harmless WoW Collector, its
      affiliates, and their respective directors, officers, employees, agents,
      and partners from and against any claims, liabilities, damages, losses,
      costs, expenses, or fees (including reasonable attorneys' fees) that arise
      from or relate to your use of the Website, violation of these Terms, or
      violation of any rights of another party.
    </p>

    <h2 class="text-2xl font-semibold mb-4">13. Changes to the Terms</h2>
    <p class="mb-6">
      We reserve the right to modify or update these Terms at any time. If we
      make material changes to these Terms, we will notify you by email or by
      posting a notice on the Website. Your continued use of the Website after
      any such changes constitutes your acceptance of the new Terms.
    </p>

    <div class="w-full">
      <UButton variant="ghost" :icon="Icons.BACK" @click="router.back()">
        Go back
      </UButton>
    </div>
  </div>
</template>
