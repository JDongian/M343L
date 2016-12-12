for (; i >= 0; i--) {
    word = bn_get_words(scalar)[i];
    while (mask) {
        BN_consttime_swap(word & mask, x1, x2, bn_get_top(group->field));
        BN_consttime_swap(word & mask, z1, z2, bn_get_top(group->field));
        if (!gf2m_Madd(group, point->X, x2, z2, x1, z1, ctx))
            goto err;
        if (!gf2m_Mdouble(group, x1, z1, ctx))
            goto err;
        BN_consttime_swap(word & mask, x1, x2, bn_get_top(group->field));
        BN_consttime_swap(word & mask, z1, z2, bn_get_top(group->field));
        mask >>= 1;
    }
    mask = BN_TBIT;
}   
